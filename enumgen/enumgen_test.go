// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"goki.dev/enums"
	"goki.dev/enums/enumgen/config"
	"goki.dev/enums/enumgen/testdata"
	"goki.dev/grease"
)

func TestGenerate(t *testing.T) {
	c := &config.Config{}
	err := grease.SetFromDefaults(c)
	if err != nil {
		t.Errorf("programmer error: error setting config from default tags: %v", err)
	}
	c.Dir = "./testdata"
	c.Output = "./testdata/enumgen.go"
	c.JSON = true
	err = Generate(c)
	if err != nil {
		t.Errorf("error while generating: %v", err)
	}
	have, err := os.ReadFile("testdata/enumgen.go")
	if err != nil {
		t.Errorf("error while reading generated file: %v", err)
	}
	want, err := os.ReadFile("testdata/enumgen.golden")
	if err != nil {
		t.Errorf("error while reading golden file: %v", err)
	}
	// ignore first line, which has "Code generated by" message
	// that can change based on where go test is ran.
	_, shave, got := strings.Cut(string(have), "\n")
	if !got {
		t.Errorf("expected string with newline in testdata/enumgen.go, but got %q", have)
	}
	_, swant, got := strings.Cut(string(want), "\n")
	if !got {
		t.Errorf("expected string with newline in testdata/enumgen.golden, but got %q", want)
	}
	if shave != swant {
		t.Errorf("expected generated file and expected file to be the same after the first line, but they are not (compare ./testdata/enumgen.go and ./testdata/enumgen.golden to see the difference)")
	}
}

func TestFruitsString(t *testing.T) {
	val := testdata.Peach
	want := "Peach"
	have := val.String()
	if have != want {
		t.Errorf("expected string value for %d to be %q but got %q", val, want, have)
	}
}

func TestFruitsSetString(t *testing.T) {
	src := "apricot"
	want := testdata.Apricot
	var have testdata.Fruits
	err := have.SetString(src)
	if err != nil {
		t.Errorf("error setting from string %q: %v", src, err)
	}
	if have != want {
		t.Errorf("expected value %v from string %q, but got %v", want, src, have)
	}
}

func TestFoodsString(t *testing.T) {
	val := testdata.Foods(testdata.Blackberry)
	want := "Blackberry"
	have := val.String()
	if have != want {
		t.Errorf("expected string value for %d to be %q but got %q", val, want, have)
	}
}

func TestFoodsSetString(t *testing.T) {
	src := "apricot"
	want := testdata.Foods(testdata.Apricot)
	var have testdata.Foods
	err := have.SetString(src)
	if err != nil {
		t.Errorf("error setting from string %q: %v", src, err)
	}
	if have != want {
		t.Errorf("expected value %v from string %q, but got %v", want, src, have)
	}
}

func TestFoodsIsValid(t *testing.T) {
	if !testdata.Foods(testdata.Blueberry).IsValid() {
		t.Errorf("expected value Blueberry to be a valid food, but it is not")
	}
	if !testdata.Meat.IsValid() {
		t.Errorf("expected value Meat to be a valid food, but it is not")
	}
}

func TestMoreLanguagesDesc(t *testing.T) {
	val := testdata.MoreLanguages(testdata.JavaScript)
	want := "JavaScript is the worst programming language"
	have := val.Desc()
	if have != want {
		t.Errorf("expected description for value %s to be %q, but got %q", val.BitIndexString(), want, have)
	}
}

func TestFoodsValuesMethod(t *testing.T) {
	want := []testdata.Foods{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa}
	have := testdata.FoodsValues()
	if fmt.Sprintf("%v", want) != fmt.Sprintf("%v", have) {
		t.Errorf("expected foods values to be %v, but got %v", want, have)
	}
}

func TestStatesStringSetString(t *testing.T) {
	var val testdata.States
	val.SetFlag(true, testdata.Active, testdata.Hovered, testdata.Focused)
	orig := val
	want := "focused|vered|currently-being-pressed-by-user"
	have := val.String()
	if have != want {
		t.Errorf("expected string value for %d to be %q but got %q", val, want, have)
	}
	err := val.SetString(have)
	if err != nil {
		t.Errorf("error setting value from string %q: %v", have, err)
	}
	if val != orig {
		t.Errorf("new value %v after going to and from string not the same as old value %v", val, orig)
	}
}

func TestStatesSetStringString(t *testing.T) {
	src := "enabled|focused|selected"
	var want testdata.States
	want.SetFlag(true, testdata.Enabled, testdata.Focused, testdata.Selected)
	var have testdata.States
	err := have.SetString(src)
	if err != nil {
		t.Errorf("error setting value from string %q: %v", src, err)
	}
	if have != want {
		t.Errorf("expected value %v from string %q, but got %v", want, src, have)
	}
	str := have.String()
	if str != src {
		t.Errorf("expected string value for %d to be %q but got %q", have, src, str)
	}
}

func TestLanguagesString(t *testing.T) {
	var val testdata.Languages
	val.SetFlag(true, testdata.Dart, testdata.Go, testdata.Kotlin, testdata.JavaScript)
	want := "Go|JavaScript|Dart|Kotlin"
	have := val.String()
	if have != want {
		t.Errorf("expected string value for %d to be %q but got %q", val, want, have)
	}
}

func TestMoreLanguagesString(t *testing.T) {
	var val testdata.MoreLanguages
	val.SetFlag(true, testdata.Go, testdata.Perl, testdata.Python, testdata.Dart)
	val.SetFlag(false, testdata.Python)
	want := "Go|Dart|Perl"
	have := val.String()
	if have != want {
		t.Errorf("expected string value for %d to be %q but got %q", val, want, have)
	}
}

func TestMoreLanguagesSetString(t *testing.T) {
	src := "Perl|JavaScript|Kotlin"
	var have testdata.MoreLanguages
	var want testdata.MoreLanguages
	want.SetFlag(true, testdata.Perl, testdata.JavaScript, testdata.Kotlin)
	err := have.SetString(src)
	if err != nil {
		t.Errorf("error setting value from string %q: %v", src, err)
	}
	if have != want {
		t.Errorf("expected value %v from string %q, but got %v", want, src, have)
	}
}

func TestMoreLanguagesValuesGlobal(t *testing.T) {
	// need to use loop to get slice of enums.Enum without typing in constant name for everything
	wantl := []testdata.MoreLanguages{6, 10, 14, 18, 22, 26, 30, 34, 38, 42, 46, 50, 54, 55}
	want := []enums.Enum{}
	for _, i := range wantl {
		want = append(want, i)
	}
	have := testdata.MoreLanguagesN.Values()
	if fmt.Sprintf("%#v", want) != fmt.Sprintf("%#v", have) {
		t.Errorf("expected more languages values to be %#v, but got %#v", want, have)
	}
}
