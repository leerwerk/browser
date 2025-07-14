// Copyright 2025 Leer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package browser

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

func getFullFilePath(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "windows":
		return path, nil
	default:
		return filepath.Join("file://", path), nil
	}
}

func openFile(path string) error {
	fullPath, err := getFullFilePath(path)
	if err != nil {
		return err
	}

	return openURL(fullPath)
}

func openFileByBrowser(browser Browser, path string) error {
	fullPath, err := getFullFilePath(path)
	if err != nil {
		return err
	}

	return openURLByBrowser(browser, fullPath)
}

func openURLByBrowser(browser Browser, url string) error {
	return execCommand(browser.Path, url)
}

func execCommand(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func detectDefaultBrowser() (browser Browser) {
	browser.Path = getDefaultBrowserPath()
	if browser.Path == "" {
		return
	}

	for _, b := range osDefaultBrowsers {
		if strings.Contains(browser.Path, strings.ToLower(b.Name)) {
			browser.Name = b.Name
			break
		}
	}

	browser.Version = getBrowserVersion(browser.Name, browser.Path)

	return
}

func detectInstalledBrowsers() []Browser {
	var installedBrowsers []Browser

	defaultBrowserPath := getDefaultBrowserPath()
	defaultIndex := 0

	for _, browser := range osDefaultBrowsers {
		if _, err := os.Stat(browser.Path); err == nil {

			installedBrowser := Browser{
				Name: browser.Name,
				Path: browser.Path,
			}

			installedBrowser.Version = getBrowserVersion(installedBrowser.Name, installedBrowser.Path)

			installedBrowsers = append(installedBrowsers, installedBrowser)

			if installedBrowser.Path == defaultBrowserPath {
				defaultIndex = len(installedBrowsers) - 1
			}
		}
	}

	if defaultIndex != 0 {
		item := installedBrowsers[defaultIndex]
		return append([]Browser{item}, append(installedBrowsers[:defaultIndex], installedBrowsers[defaultIndex+1:]...)...)
	}

	return installedBrowsers
}

func getDigits(str string) string {
	// get the version from "Google Chrome 138.0.7204.100"
	fields := strings.Fields(str)
	for _, field := range fields {
		// check if the field starts with a digit
		if len(field) > 0 && unicode.IsDigit(rune(field[0])) {
			return strings.TrimSpace(field)
		}
	}

	return strings.TrimSpace(str)
}
