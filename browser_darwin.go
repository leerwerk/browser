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
	"bytes"
	"os"
	"os/exec"
	"strings"
)

var osDefaultBrowsers = []Browser{
	{Name: Safari, Path: `/Applications/Safari.app/Contents/MacOS/Safari`},
	{Name: Chrome, Path: `/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`},
	{Name: Firefox, Path: `/Applications/Firefox.app/Contents/MacOS/firefox`},
	{Name: Chromium, Path: `/Applications/Chromium.app/Contents/MacOS/Chromium`},
	{Name: Edge, Path: `/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge`},
}

func openURL(url string) error {
	return execCommand("open", url)
}

func getDefaultBrowserPath() string {
	var out bytes.Buffer

	cmd := exec.Command("defaults", "read", "com.apple.LaunchServices/com.apple.launchservices.secure", "LSHandlers")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return ""
	}

	return strings.ToLower(strings.TrimSpace(out.String()))
}

func getBrowserVersion(name, path string) string {
	if name == Safari {
		return _getSafariVersion()
	}

	var out bytes.Buffer

	cmd := exec.Command(path, "--version")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return Unknown
	}

	return getDigits(out.String())
}

func _getSafariVersion() string {
	const plistPath = "/Applications/Safari.app/Contents/Info.plist"

	// Check if file exists
	if _, err := os.Stat(plistPath); os.IsNotExist(err) {
		return Unknown
	}

	var out bytes.Buffer

	// Use plutil to extract the version string
	cmd := exec.Command("plutil", "-extract", "CFBundleShortVersionString", "raw", "-o", "-", plistPath)
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return Unknown
	}

	return getDigits(out.String())
}
