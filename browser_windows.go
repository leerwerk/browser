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
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var osDefaultBrowsers = []Browser{
	{Name: Edge, Path: `C:\Program Files\Microsoft\Edge\Application\msedge.exe`},
	{Name: Edge, Path: `C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`},
	{Name: Chrome, Path: `C:\Program Files\Google\Chrome\Application\chrome.exe`},
	{Name: Chrome, Path: `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`},
	{Name: Chromium, Path: `C:\Program Files\Chromium\Application\chrome.exe`},
	{Name: Chromium, Path: `C:\Program Files (x86)\Chromium\Application\chrome.exe`},
	{Name: Firefox, Path: `C:\Program Files\Mozilla Firefox\firefox.exe`},
	{Name: Firefox, Path: `C:\Program Files (x86)\Mozilla Firefox\firefox.exe`},
}

func openURL(url string) error {
	return execCommand("rundll32", "url.dll,FileProtocolHandler", url)
}

var _reg = regexp.MustCompile(`"([^"]+\.exe)"`)

func getDefaultBrowserPath() string {
	var out bytes.Buffer

	cmd := exec.Command("reg", "query", `HKEY_CLASSES_ROOT\http\shell\open\command`, "/ve")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return ""
	}

	// "\r\nHKEY_CLASSES_ROOT\\http\\shell\\open\\command\r\n    (Default)    REG_SZ    \"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe\" \"%1\"\r\n\r\n"
	// Extract the quoted path using regex
	matches := _reg.FindStringSubmatch(out.String())
	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func getBrowserVersion(_, path string) string {
	var out bytes.Buffer

	cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`(Get-Item '%s').VersionInfo.ProductVersion`, path))
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return Unknown
	}

	return strings.TrimSpace(out.String())
}
