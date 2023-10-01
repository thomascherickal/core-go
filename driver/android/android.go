// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on https://github.com/golang/mobile
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build android

/*
Android Apps are built with -buildmode=c-shared. They are loaded by a
running Java process.

Before any entry point is reached, a global constructor initializes the
Go runtime, calling all Go init functions. All cgo calls will block
until this is complete. Next JNI_OnLoad is called. When that is
complete, one of two entry points is called.

All-Go apps built using NativeActivity enter at ANativeActivity_onCreate.
*/
package android

/*
#cgo LDFLAGS: -landroid -llog

#include <android/configuration.h>
#include <android/input.h>
#include <android/keycodes.h>
#include <android/looper.h>
#include <android/native_activity.h>
#include <android/native_window.h>
#include <jni.h>
#include <pthread.h>
#include <stdlib.h>
#include <stdbool.h>

int32_t getKeyRune(JNIEnv* env, AInputEvent* e);

void showKeyboard(JNIEnv* env, int keyboardType);
void hideKeyboard(JNIEnv* env);
void showFileOpen(JNIEnv* env, char* mimes);
void showFileSave(JNIEnv* env, char* mimes, char* filename);

void Java_org_golang_app_GoNativeActivity_filePickerReturned(JNIEnv *env, jclass clazz, jstring str);
*/
import "C"
import (
	"fmt"
	"image"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"
	"unsafe"

	"goki.dev/goosi"
	"goki.dev/goosi/driver/mobile/callfn"
	"goki.dev/goosi/driver/mobile/mobileinit"
	"goki.dev/goosi/events"
	"goki.dev/goosi/events/key"
	"goki.dev/mobile/event/size"
)

// mimeMap contains standard mime entries that are missing on Android
var mimeMap = map[string]string{
	".txt": "text/plain",
}

// RunOnJVM runs fn on a new goroutine locked to an OS thread with a JNIEnv.
//
// RunOnJVM blocks until the call to fn is complete. Any Java
// exception or failure to attach to the JVM is returned as an error.
//
// The function fn takes vm, the current JavaVM*,
// env, the current JNIEnv*, and
// ctx, a jobject representing the global android.context.Context.
func RunOnJVM(fn func(vm, jniEnv, ctx uintptr) error) error {
	return mobileinit.RunOnJVM(fn)
}

//export setCurrentContext
func setCurrentContext(vm *C.JavaVM, ctx C.jobject) {
	mobileinit.SetCurrentContext(unsafe.Pointer(vm), uintptr(ctx))
}

//export callMain
func callMain(mainPC uintptr) {
	fmt.Println("calling main")
	for _, name := range []string{"FILESDIR", "TMPDIR", "PATH", "LD_LIBRARY_PATH"} {
		n := C.CString(name)
		os.Setenv(name, C.GoString(C.getenv(n)))
		C.free(unsafe.Pointer(n))
	}

	// Set timezone.
	//
	// Note that Android zoneinfo is stored in /system/usr/share/zoneinfo,
	// but it is in some kind of packed TZiff file that we do not support
	// yet. As a stopgap, we build a fixed zone using the tm_zone name.
	var curtime C.time_t
	var curtm C.struct_tm
	C.time(&curtime)
	C.localtime_r(&curtime, &curtm)
	tzOffset := int(curtm.tm_gmtoff)
	tz := C.GoString(curtm.tm_zone)
	time.Local = time.FixedZone(tz, tzOffset)

	go callfn.CallFn(mainPC)
}

//export onStart
func onStart(activity *C.ANativeActivity) {
	fmt.Println("started")
}

//export onResume
func onResume(activity *C.ANativeActivity) {
}

//export onSaveInstanceState
func onSaveInstanceState(activity *C.ANativeActivity, outSize *C.size_t) unsafe.Pointer {
	return nil
}

//export onPause
func onPause(activity *C.ANativeActivity) {
}

//export onStop
func onStop(activity *C.ANativeActivity) {
}

//export onCreate
func onCreate(activity *C.ANativeActivity) {
	fmt.Println("created")
	// Set the initial configuration.
	//
	// Note we use unbuffered channels to talk to the activity loop, and
	// NativeActivity calls these callbacks sequentially, so configuration
	// will be set before <-windowRedrawNeeded is processed.
	windowConfigChange <- windowConfigRead(activity)
}

//export onDestroy
func onDestroy(activity *C.ANativeActivity) {
	activityDestroyed <- struct{}{}
}

//export onWindowFocusChanged
func onWindowFocusChanged(activity *C.ANativeActivity, hasFocus C.int) {
}

//export onNativeWindowCreated
func onNativeWindowCreated(activity *C.ANativeActivity, window *C.ANativeWindow) {
	theApp.winptr = uintptr(unsafe.Pointer(window))
	fmt.Println("win creATED", theApp.winptr)
	theApp.setSysWindow(nil, theApp.winptr)
}

//export onNativeWindowRedrawNeeded
func onNativeWindowRedrawNeeded(activity *C.ANativeActivity, window *C.ANativeWindow) {
	// Called on orientation change and window resize.
	// Send a request for redraw, and block this function
	// until a complete draw and buffer swap is completed.
	// This is required by the redraw documentation to
	// avoid bad draws.
	windowRedrawNeeded <- window
	<-windowRedrawDone
}

//export onNativeWindowDestroyed
func onNativeWindowDestroyed(activity *C.ANativeActivity, window *C.ANativeWindow) {
	windowDestroyed <- window
}

//export onInputQueueCreated
func onInputQueueCreated(activity *C.ANativeActivity, q *C.AInputQueue) {
	inputQueue <- q
	<-inputQueueDone
}

//export onInputQueueDestroyed
func onInputQueueDestroyed(activity *C.ANativeActivity, q *C.AInputQueue) {
	inputQueue <- nil
	<-inputQueueDone
}

//export onContentRectChanged
func onContentRectChanged(activity *C.ANativeActivity, rect *C.ARect) {
}

//export setDarkMode
func setDarkMode(dark C.bool) {
	theApp.darkMode = bool(dark)
}

type windowConfig struct {
	orientation size.Orientation
	dotsPerPx   float32 // raw display dots per standard pixel (1/96 of 1 in)
}

func windowConfigRead(activity *C.ANativeActivity) windowConfig {
	aconfig := C.AConfiguration_new()
	C.AConfiguration_fromAssetManager(aconfig, activity.assetManager)
	orient := C.AConfiguration_getOrientation(aconfig)
	density := C.AConfiguration_getDensity(aconfig)
	C.AConfiguration_delete(aconfig)

	// Calculate the screen resolution. This value is approximate. For example,
	// a physical resolution of 200 DPI may be quantized to one of the
	// ACONFIGURATION_DENSITY_XXX values such as 160 or 240.
	//
	// A more accurate DPI could possibly be calculated from
	// https://developer.android.com/reference/android/util/DisplayMetrics.html#xdpi
	// but this does not appear to be accessible via the NDK. In any case, the
	// hardware might not even provide a more accurate number, as the system
	// does not apparently use the reported value. See golang.org/issue/13366
	// for a discussion.
	var dpi int
	switch density {
	case C.ACONFIGURATION_DENSITY_DEFAULT:
		dpi = 160
	case C.ACONFIGURATION_DENSITY_LOW,
		C.ACONFIGURATION_DENSITY_MEDIUM,
		213, // C.ACONFIGURATION_DENSITY_TV
		C.ACONFIGURATION_DENSITY_HIGH,
		320, // ACONFIGURATION_DENSITY_XHIGH
		480, // ACONFIGURATION_DENSITY_XXHIGH
		640: // ACONFIGURATION_DENSITY_XXXHIGH
		dpi = int(density)
	case C.ACONFIGURATION_DENSITY_NONE:
		log.Print("android device reports no screen density")
		dpi = 72
	default:
		// TODO: fix this always happening with value 240
		log.Printf("android device reports unknown density: %d", density)
		// All we can do is guess.
		if density > 0 {
			dpi = int(density)
		} else {
			dpi = 72
		}
	}

	o := size.OrientationUnknown
	switch orient {
	case C.ACONFIGURATION_ORIENTATION_PORT:
		o = size.OrientationPortrait
	case C.ACONFIGURATION_ORIENTATION_LAND:
		o = size.OrientationLandscape
	}

	return windowConfig{
		orientation: o,
		dotsPerPx:   float32(dpi) / 96,
	}
}

//export onConfigurationChanged
func onConfigurationChanged(activity *C.ANativeActivity) {
	// A rotation event first triggers onConfigurationChanged, then
	// calls onNativeWindowRedrawNeeded. We extract the orientation
	// here and save it for the redraw event.
	windowConfigChange <- windowConfigRead(activity)
}

//export onLowMemory
func onLowMemory(activity *C.ANativeActivity) {
	runtime.GC()
	debug.FreeOSMemory()
}

var (
	inputQueue         = make(chan *C.AInputQueue)
	inputQueueDone     = make(chan struct{})
	windowDestroyed    = make(chan *C.ANativeWindow)
	windowRedrawNeeded = make(chan *C.ANativeWindow)
	windowRedrawDone   = make(chan struct{})
	windowConfigChange = make(chan windowConfig)
	activityDestroyed  = make(chan struct{})
)

func main(f func(*appImpl)) {
	fmt.Println("in main")
	mainUserFn = f
	// TODO: merge the runInputQueue and mainUI functions?
	go func() {
		fmt.Println("running input queue")
		if err := mobileinit.RunOnJVM(runInputQueue); err != nil {
			log.Fatalf("app: %v", err)
		}
	}()
	// Preserve this OS thread for:
	//	1. the attached JNI thread
	fmt.Println("running main UI")
	if err := mobileinit.RunOnJVM(theApp.mainUI); err != nil {
		log.Fatalf("app: %v", err)
	}
}

// ShowVirtualKeyboard requests the driver to show a virtual keyboard for text input
func (a *appImpl) ShowVirtualKeyboard(typ goosi.VirtualKeyboardTypes) {
	err := mobileinit.RunOnJVM(func(vm, jniEnv, ctx uintptr) error {
		env := (*C.JNIEnv)(unsafe.Pointer(jniEnv)) // not a Go heap pointer
		C.showKeyboard(env, C.int(int32(typ)))
		return nil
	})
	if err != nil {
		log.Fatalf("app: %v", err)
	}
}

// HideVirtualKeyboard requests the driver to hide any visible virtual keyboard
func (a *appImpl) HideVirtualKeyboard() {
	if err := mobileinit.RunOnJVM(hideSoftInput); err != nil {
		log.Fatalf("app: %v", err)
	}
}

func hideSoftInput(vm, jniEnv, ctx uintptr) error {
	env := (*C.JNIEnv)(unsafe.Pointer(jniEnv)) // not a Go heap pointer
	C.hideKeyboard(env)
	return nil
}

//export insetsChanged
func insetsChanged(top, bottom, left, right int) {
	theApp.insets.Set(float32(top), float32(right), float32(bottom), float32(left))
}

var mainUserFn func(*appImpl)

func (app *appImpl) mainUI(vm, jniEnv, ctx uintptr) error {
	donec := make(chan struct{})
	go func() {
		mainUserFn(theApp)
		close(donec)
	}()

	var dotsPerPx float32

	for {
		select {
		case <-donec:
			return nil
		case cfg := <-windowConfigChange:
			dotsPerPx = cfg.dotsPerPx
		case w := <-windowRedrawNeeded:
			app.window.EvMgr.Window(events.Focus)

			widthDots := int(C.ANativeWindow_getWidth(w))
			heightDots := int(C.ANativeWindow_getHeight(w))

			app.screen.ScreenNumber = 0
			app.screen.DevicePixelRatio = dotsPerPx
			wsz := image.Point{widthDots, heightDots}
			app.screen.Geometry = image.Rectangle{Max: wsz}
			app.screen.PixSize = app.screen.WinSizeToPix(wsz)
			app.screen.Orientation = screenOrientation(widthDots, heightDots)
			app.screen.UpdatePhysicalDPI()
			app.screen.UpdateLogicalDPI()

			app.window.PhysDPI = app.screen.PhysicalDPI
			app.window.PxSize = app.screen.PixSize
			app.window.WnSize = wsz

			app.window.EvMgr.WindowPaint()
		case <-windowDestroyed:
			app.window.EvMgr.Window(events.Show) // TODO: does this make sense? it is based on the gomobile code
		case <-activityDestroyed:
			app.window.EvMgr.Window(events.Close)
			// case <-app.publish: // TODO(kai): do something here?
			// 	select {
			// 	case windowRedrawDone <- struct{}{}:
			// 	default:
			// 	}
		}
	}
}

func screenOrientation(width, height int) goosi.ScreenOrientation {
	if width > height {
		return goosi.Landscape
	}
	return goosi.Portrait
}

func runInputQueue(vm, jniEnv, ctx uintptr) error {
	env := (*C.JNIEnv)(unsafe.Pointer(jniEnv)) // not a Go heap pointer

	// Android loopers select on OS file descriptors, not Go channels, so we
	// translate the inputQueue channel to an ALooper_wake call.
	l := C.ALooper_prepare(C.ALOOPER_PREPARE_ALLOW_NON_CALLBACKS)
	pending := make(chan *C.AInputQueue, 1)
	go func() {
		for q := range inputQueue {
			pending <- q
			C.ALooper_wake(l)
		}
	}()

	var q *C.AInputQueue
	for {
		if C.ALooper_pollAll(-1, nil, nil, nil) == C.ALOOPER_POLL_WAKE {
			select {
			default:
			case p := <-pending:
				if q != nil {
					processEvents(env, q)
					C.AInputQueue_detachLooper(q)
				}
				q = p
				if q != nil {
					C.AInputQueue_attachLooper(q, l, 0, nil, nil)
				}
				inputQueueDone <- struct{}{}
			}
		}
		if q != nil {
			processEvents(env, q)
		}
	}
}

func processEvents(env *C.JNIEnv, q *C.AInputQueue) {
	var e *C.AInputEvent
	for C.AInputQueue_getEvent(q, &e) >= 0 {
		if C.AInputQueue_preDispatchEvent(q, e) != 0 {
			continue
		}
		processEvent(env, e)
		C.AInputQueue_finishEvent(q, e, 0)
	}
}

func processEvent(env *C.JNIEnv, e *C.AInputEvent) {
	switch C.AInputEvent_getType(e) {
	case C.AINPUT_EVENT_TYPE_KEY:
		processKey(env, e)
	case C.AINPUT_EVENT_TYPE_MOTION:
		// At most one of the events in this batch is an up or down event; get its index and change.
		upDownIndex := C.size_t(C.AMotionEvent_getAction(e)&C.AMOTION_EVENT_ACTION_POINTER_INDEX_MASK) >> C.AMOTION_EVENT_ACTION_POINTER_INDEX_SHIFT
		upDownType := events.TouchMove
		switch C.AMotionEvent_getAction(e) & C.AMOTION_EVENT_ACTION_MASK {
		case C.AMOTION_EVENT_ACTION_DOWN, C.AMOTION_EVENT_ACTION_POINTER_DOWN:
			upDownType = events.TouchStart
		case C.AMOTION_EVENT_ACTION_UP, C.AMOTION_EVENT_ACTION_POINTER_UP:
			upDownType = events.TouchEnd
		}

		for i, n := C.size_t(0), C.AMotionEvent_getPointerCount(e); i < n; i++ {
			t := events.TouchMove
			if i == upDownIndex {
				t = upDownType
			}
			seq := events.Sequence(C.AMotionEvent_getPointerId(e, i))
			x := int(C.AMotionEvent_getX(e, i))
			y := int(C.AMotionEvent_getY(e, i))
			theApp.window.EvMgr.Touch(t, seq, image.Pt(x, y))
		}
	default:
		log.Printf("unknown input event, type=%d", C.AInputEvent_getType(e))
	}
}

func processKey(env *C.JNIEnv, e *C.AInputEvent) {
	deviceID := C.AInputEvent_getDeviceId(e)
	if deviceID == 0 {
		// Software keyboard input, leaving for scribe/IME.
		return
	}

	r := rune(C.getKeyRune(env, e))
	code := convAndroidKeyCode(int32(C.AKeyEvent_getKeyCode(e)))

	if r >= '0' && r <= '9' { // GBoard generates key events for numbers, but we see them in textChanged
		return
	}
	typ := events.KeyDown
	if C.AKeyEvent_getAction(e) == C.AKEY_STATE_UP {
		typ = events.KeyUp
	}
	// TODO(crawshaw): set Modifiers.
	theApp.window.EvMgr.Key(typ, r, code, 0)
}

var androidKeycodes = map[int32]key.Codes{
	C.AKEYCODE_HOME:            key.CodeHome,
	C.AKEYCODE_0:               key.Code0,
	C.AKEYCODE_1:               key.Code1,
	C.AKEYCODE_2:               key.Code2,
	C.AKEYCODE_3:               key.Code3,
	C.AKEYCODE_4:               key.Code4,
	C.AKEYCODE_5:               key.Code5,
	C.AKEYCODE_6:               key.Code6,
	C.AKEYCODE_7:               key.Code7,
	C.AKEYCODE_8:               key.Code8,
	C.AKEYCODE_9:               key.Code9,
	C.AKEYCODE_VOLUME_UP:       key.CodeVolumeUp,
	C.AKEYCODE_VOLUME_DOWN:     key.CodeVolumeDown,
	C.AKEYCODE_A:               key.CodeA,
	C.AKEYCODE_B:               key.CodeB,
	C.AKEYCODE_C:               key.CodeC,
	C.AKEYCODE_D:               key.CodeD,
	C.AKEYCODE_E:               key.CodeE,
	C.AKEYCODE_F:               key.CodeF,
	C.AKEYCODE_G:               key.CodeG,
	C.AKEYCODE_H:               key.CodeH,
	C.AKEYCODE_I:               key.CodeI,
	C.AKEYCODE_J:               key.CodeJ,
	C.AKEYCODE_K:               key.CodeK,
	C.AKEYCODE_L:               key.CodeL,
	C.AKEYCODE_M:               key.CodeM,
	C.AKEYCODE_N:               key.CodeN,
	C.AKEYCODE_O:               key.CodeO,
	C.AKEYCODE_P:               key.CodeP,
	C.AKEYCODE_Q:               key.CodeQ,
	C.AKEYCODE_R:               key.CodeR,
	C.AKEYCODE_S:               key.CodeS,
	C.AKEYCODE_T:               key.CodeT,
	C.AKEYCODE_U:               key.CodeU,
	C.AKEYCODE_V:               key.CodeV,
	C.AKEYCODE_W:               key.CodeW,
	C.AKEYCODE_X:               key.CodeX,
	C.AKEYCODE_Y:               key.CodeY,
	C.AKEYCODE_Z:               key.CodeZ,
	C.AKEYCODE_COMMA:           key.CodeComma,
	C.AKEYCODE_PERIOD:          key.CodeFullStop,
	C.AKEYCODE_ALT_LEFT:        key.CodeLeftAlt,
	C.AKEYCODE_ALT_RIGHT:       key.CodeRightAlt,
	C.AKEYCODE_SHIFT_LEFT:      key.CodeLeftShift,
	C.AKEYCODE_SHIFT_RIGHT:     key.CodeRightShift,
	C.AKEYCODE_TAB:             key.CodeTab,
	C.AKEYCODE_SPACE:           key.CodeSpacebar,
	C.AKEYCODE_ENTER:           key.CodeReturnEnter,
	C.AKEYCODE_DEL:             key.CodeDeleteBackspace,
	C.AKEYCODE_GRAVE:           key.CodeGraveAccent,
	C.AKEYCODE_MINUS:           key.CodeHyphenMinus,
	C.AKEYCODE_EQUALS:          key.CodeEqualSign,
	C.AKEYCODE_LEFT_BRACKET:    key.CodeLeftSquareBracket,
	C.AKEYCODE_RIGHT_BRACKET:   key.CodeRightSquareBracket,
	C.AKEYCODE_BACKSLASH:       key.CodeBackslash,
	C.AKEYCODE_SEMICOLON:       key.CodeSemicolon,
	C.AKEYCODE_APOSTROPHE:      key.CodeApostrophe,
	C.AKEYCODE_SLASH:           key.CodeSlash,
	C.AKEYCODE_PAGE_UP:         key.CodePageUp,
	C.AKEYCODE_PAGE_DOWN:       key.CodePageDown,
	C.AKEYCODE_ESCAPE:          key.CodeEscape,
	C.AKEYCODE_FORWARD_DEL:     key.CodeDeleteForward,
	C.AKEYCODE_CTRL_LEFT:       key.CodeLeftControl,
	C.AKEYCODE_CTRL_RIGHT:      key.CodeRightControl,
	C.AKEYCODE_CAPS_LOCK:       key.CodeCapsLock,
	C.AKEYCODE_META_LEFT:       key.CodeLeftMeta,
	C.AKEYCODE_META_RIGHT:      key.CodeRightMeta,
	C.AKEYCODE_INSERT:          key.CodeInsert,
	C.AKEYCODE_F1:              key.CodeF1,
	C.AKEYCODE_F2:              key.CodeF2,
	C.AKEYCODE_F3:              key.CodeF3,
	C.AKEYCODE_F4:              key.CodeF4,
	C.AKEYCODE_F5:              key.CodeF5,
	C.AKEYCODE_F6:              key.CodeF6,
	C.AKEYCODE_F7:              key.CodeF7,
	C.AKEYCODE_F8:              key.CodeF8,
	C.AKEYCODE_F9:              key.CodeF9,
	C.AKEYCODE_F10:             key.CodeF10,
	C.AKEYCODE_F11:             key.CodeF11,
	C.AKEYCODE_F12:             key.CodeF12,
	C.AKEYCODE_NUM_LOCK:        key.CodeKeypadNumLock,
	C.AKEYCODE_NUMPAD_0:        key.CodeKeypad0,
	C.AKEYCODE_NUMPAD_1:        key.CodeKeypad1,
	C.AKEYCODE_NUMPAD_2:        key.CodeKeypad2,
	C.AKEYCODE_NUMPAD_3:        key.CodeKeypad3,
	C.AKEYCODE_NUMPAD_4:        key.CodeKeypad4,
	C.AKEYCODE_NUMPAD_5:        key.CodeKeypad5,
	C.AKEYCODE_NUMPAD_6:        key.CodeKeypad6,
	C.AKEYCODE_NUMPAD_7:        key.CodeKeypad7,
	C.AKEYCODE_NUMPAD_8:        key.CodeKeypad8,
	C.AKEYCODE_NUMPAD_9:        key.CodeKeypad9,
	C.AKEYCODE_NUMPAD_DIVIDE:   key.CodeKeypadSlash,
	C.AKEYCODE_NUMPAD_MULTIPLY: key.CodeKeypadAsterisk,
	C.AKEYCODE_NUMPAD_SUBTRACT: key.CodeKeypadHyphenMinus,
	C.AKEYCODE_NUMPAD_ADD:      key.CodeKeypadPlusSign,
	C.AKEYCODE_NUMPAD_DOT:      key.CodeKeypadFullStop,
	C.AKEYCODE_NUMPAD_ENTER:    key.CodeKeypadEnter,
	C.AKEYCODE_NUMPAD_EQUALS:   key.CodeKeypadEqualSign,
	C.AKEYCODE_VOLUME_MUTE:     key.CodeMute,
}

func convAndroidKeyCode(aKeyCode int32) key.Codes {
	if code, ok := androidKeycodes[aKeyCode]; ok {
		return code
	}
	return key.CodeUnknown
}

/*
	Many Android key codes do not map into USB HID codes.
	For those, key.CodeUnknown is returned. This switch has all
	cases, even the unknown ones, to serve as a documentation
	and search aid.
	C.AKEYCODE_UNKNOWN
	C.AKEYCODE_SOFT_LEFT
	C.AKEYCODE_SOFT_RIGHT
	C.AKEYCODE_BACK
	C.AKEYCODE_CALL
	C.AKEYCODE_ENDCALL
	C.AKEYCODE_STAR
	C.AKEYCODE_POUND
	C.AKEYCODE_DPAD_UP
	C.AKEYCODE_DPAD_DOWN
	C.AKEYCODE_DPAD_LEFT
	C.AKEYCODE_DPAD_RIGHT
	C.AKEYCODE_DPAD_CENTER
	C.AKEYCODE_POWER
	C.AKEYCODE_CAMERA
	C.AKEYCODE_CLEAR
	C.AKEYCODE_SYM
	C.AKEYCODE_EXPLORER
	C.AKEYCODE_ENVELOPE
	C.AKEYCODE_AT
	C.AKEYCODE_NUM
	C.AKEYCODE_HEADSETHOOK
	C.AKEYCODE_FOCUS
	C.AKEYCODE_PLUS
	C.AKEYCODE_MENU
	C.AKEYCODE_NOTIFICATION
	C.AKEYCODE_SEARCH
	C.AKEYCODE_MEDIA_PLAY_PAUSE
	C.AKEYCODE_MEDIA_STOP
	C.AKEYCODE_MEDIA_NEXT
	C.AKEYCODE_MEDIA_PREVIOUS
	C.AKEYCODE_MEDIA_REWIND
	C.AKEYCODE_MEDIA_FAST_FORWARD
	C.AKEYCODE_MUTE
	C.AKEYCODE_PICTSYMBOLS
	C.AKEYCODE_SWITCH_CHARSET
	C.AKEYCODE_BUTTON_A
	C.AKEYCODE_BUTTON_B
	C.AKEYCODE_BUTTON_C
	C.AKEYCODE_BUTTON_X
	C.AKEYCODE_BUTTON_Y
	C.AKEYCODE_BUTTON_Z
	C.AKEYCODE_BUTTON_L1
	C.AKEYCODE_BUTTON_R1
	C.AKEYCODE_BUTTON_L2
	C.AKEYCODE_BUTTON_R2
	C.AKEYCODE_BUTTON_THUMBL
	C.AKEYCODE_BUTTON_THUMBR
	C.AKEYCODE_BUTTON_START
	C.AKEYCODE_BUTTON_SELECT
	C.AKEYCODE_BUTTON_MODE
	C.AKEYCODE_SCROLL_LOCK
	C.AKEYCODE_FUNCTION
	C.AKEYCODE_SYSRQ
	C.AKEYCODE_BREAK
	C.AKEYCODE_MOVE_HOME
	C.AKEYCODE_MOVE_END
	C.AKEYCODE_FORWARD
	C.AKEYCODE_MEDIA_PLAY
	C.AKEYCODE_MEDIA_PAUSE
	C.AKEYCODE_MEDIA_CLOSE
	C.AKEYCODE_MEDIA_EJECT
	C.AKEYCODE_MEDIA_RECORD
	C.AKEYCODE_NUMPAD_COMMA
	C.AKEYCODE_NUMPAD_LEFT_PAREN
	C.AKEYCODE_NUMPAD_RIGHT_PAREN
	C.AKEYCODE_INFO
	C.AKEYCODE_CHANNEL_UP
	C.AKEYCODE_CHANNEL_DOWN
	C.AKEYCODE_ZOOM_IN
	C.AKEYCODE_ZOOM_OUT
	C.AKEYCODE_TV
	C.AKEYCODE_WINDOW
	C.AKEYCODE_GUIDE
	C.AKEYCODE_DVR
	C.AKEYCODE_BOOKMARK
	C.AKEYCODE_CAPTIONS
	C.AKEYCODE_SETTINGS
	C.AKEYCODE_TV_POWER
	C.AKEYCODE_TV_INPUT
	C.AKEYCODE_STB_POWER
	C.AKEYCODE_STB_INPUT
	C.AKEYCODE_AVR_POWER
	C.AKEYCODE_AVR_INPUT
	C.AKEYCODE_PROG_RED
	C.AKEYCODE_PROG_GREEN
	C.AKEYCODE_PROG_YELLOW
	C.AKEYCODE_PROG_BLUE
	C.AKEYCODE_APP_SWITCH
	C.AKEYCODE_BUTTON_1
	C.AKEYCODE_BUTTON_2
	C.AKEYCODE_BUTTON_3
	C.AKEYCODE_BUTTON_4
	C.AKEYCODE_BUTTON_5
	C.AKEYCODE_BUTTON_6
	C.AKEYCODE_BUTTON_7
	C.AKEYCODE_BUTTON_8
	C.AKEYCODE_BUTTON_9
	C.AKEYCODE_BUTTON_10
	C.AKEYCODE_BUTTON_11
	C.AKEYCODE_BUTTON_12
	C.AKEYCODE_BUTTON_13
	C.AKEYCODE_BUTTON_14
	C.AKEYCODE_BUTTON_15
	C.AKEYCODE_BUTTON_16
	C.AKEYCODE_LANGUAGE_SWITCH
	C.AKEYCODE_MANNER_MODE
	C.AKEYCODE_3D_MODE
	C.AKEYCODE_CONTACTS
	C.AKEYCODE_CALENDAR
	C.AKEYCODE_MUSIC
	C.AKEYCODE_CALCULATOR

	Defined in an NDK API version beyond what we use today:
	C.AKEYCODE_ASSIST
	C.AKEYCODE_BRIGHTNESS_DOWN
	C.AKEYCODE_BRIGHTNESS_UP
	C.AKEYCODE_RO
	C.AKEYCODE_YEN
	C.AKEYCODE_ZENKAKU_HANKAKU
*/
