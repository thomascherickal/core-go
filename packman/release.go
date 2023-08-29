// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"goki.dev/goki/config"
)

// Release releases the config project
// by calling [ReleaseApp] if it is an app
// and [ReleaseLibrary] if it is a library.
func Release(c *config.Config) error {
	if c.Type == config.TypeApp {
		return ReleaseApp(c)
	}
	return ReleaseLibrary(c)
}

// ReleaseApp releases the config app.
func ReleaseApp(c *config.Config) error {
	// TODO: implement
	return nil
}

// ReleaseLibrary releases the config library.
func ReleaseLibrary(c *config.Config) error {
	str, err := VersionFileString(c)
	if err != nil {
		return fmt.Errorf("error generating version file string: %w", err)
	}
	err = os.WriteFile(c.Release.VersionFile, []byte(str), 0666)
	if err != nil {
		return fmt.Errorf("error writing version string to version file: %w", err)
	}
	err = PushGitRelease(c)
	if err != nil {
		return fmt.Errorf("error pushing git release: %w", err)
	}
	return nil
}

// VersionFileString returns the version file string
// for a project with the given config info.
func VersionFileString(c *config.Config) (string, error) {
	var b strings.Builder
	b.WriteString("// Code generated by \"goki " + ArgsString(os.Args[1:]) + "\"; DO NOT EDIT.\n\n")
	b.WriteString("package " + c.Name + "\n\n")
	b.WriteString("const (\n")
	b.WriteString("\tVersion = \"" + c.Version + "\"\n")

	gc := exec.Command("git", "rev-parse", "--short", "HEAD")
	res, err := gc.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error getting previous git commit: %w (%s)", err, res)
	}
	b.WriteString("\tGitCommit = \"" + strings.TrimSuffix(string(res), "\n") + "\" // the commit just before the release\n")

	date := time.Now().UTC().Format("2006-01-02 15:04")
	b.WriteString("\tVersionDate = \"" + date + "\" // the date-time of the release in UTC (in the format 'YYYY-MM-DD HH:MM', which is the Go format '2006-01-02 15:04')\n")
	b.WriteString(")\n\n")
	return b.String(), nil
}

// PushGitRelease commits a release commit using Git,
// adds a version tag, and pushes the code and tags
// based on the given config info.
func PushGitRelease(c *config.Config) error {
	ac := exec.Command("git", "add", c.Release.VersionFile)
	_, err := RunCmd(ac)
	if err != nil {
		return fmt.Errorf("error adding version file: %w", err)
	}

	cc := exec.Command("git", "commit", "-am", c.Version+" release; "+c.Release.VersionFile+" updated")
	_, err = RunCmd(cc)
	if err != nil {
		return fmt.Errorf("error commiting release commit: %w", err)
	}

	tc := exec.Command("git", "tag", "-a", c.Version, "-m", c.Version+" release")
	_, err = RunCmd(tc)
	if err != nil {
		return fmt.Errorf("error tagging release: %w", err)
	}

	pc := exec.Command("git", "push")
	_, err = RunCmd(pc)
	if err != nil {
		return fmt.Errorf("error pushing commit: %w", err)
	}

	ptc := exec.Command("git", "push", "origin", "--tags")
	_, err = RunCmd(ptc)
	if err != nil {
		return fmt.Errorf("error pushing tags: %w", err)
	}

	return nil
}
