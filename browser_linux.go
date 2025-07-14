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
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var osDefaultBrowsers = []Browser{
	{Name: Firefox, Path: `/usr/bin/firefox`},
	{Name: Chrome, Path: `/usr/bin/google-chrome`},
	{Name: Chrome, Path: `/usr/bin/google-chrome-stable`},
	{Name: Chrome, Path: `/opt/google/chrome/chrome`},
	{Name: Chromium, Path: `/usr/bin/chromium`},
	{Name: Chromium, Path: `/usr/bin/chromium-browser`},
	{Name: Edge, Path: `/usr/bin/microsoft-edge`},
	{Name: Edge, Path: `/opt/microsoft/msedge/msedge`},
}

func openURL(url string) error {
	// xdg-open: part of xdg-utils package (installed by default in most modern desktop Linux distros)
	// x-www-browser: part of Debian’s alternatives system
	// www-browser: part of Debian alternatives system, text-based browsers
	progs := []string{"xdg-open", "x-www-browser", "www-browser"}

	for _, prog := range progs {
		if _, err := exec.LookPath(prog); err == nil {
			return execCommand(prog, url)
		}
	}

	return &exec.Error{Name: strings.Join(progs, ","), Err: exec.ErrNotFound}
}

func getDefaultBrowserPath() string {
	getRealPath := func(str string) string {
		defaultBrowserPath := strings.ToLower(strings.TrimSpace(str))
		if strings.HasSuffix(defaultBrowserPath, ".desktop") {
			return _resolveDesktopEntry(defaultBrowserPath)
		}
		return defaultBrowserPath
	}

	var out bytes.Buffer

	// Try xdg-mime
	cmd := exec.Command("xdg-mime", "query", "default", "x-scheme-handler/http")
	cmd.Stdout = &out
	if err := cmd.Run(); err == nil {
		return getRealPath(out.String())
	}

	out.Reset()

	// Try xdg-settings
	cmd = exec.Command("xdg-settings", "get", "default-web-browser")
	cmd.Stdout = &out
	if err := cmd.Run(); err == nil {
		return getRealPath(out.String())
	}

	out.Reset()

	// Try GNOME gsettings
	cmd = exec.Command("gsettings", "get", "org.gnome.desktop.default-applications.browser", "exec")
	cmd.Stdout = &out
	if err := cmd.Run(); err == nil {
		return getRealPath(out.String())
	}

	return ""
}

func _resolveDesktopEntry(desktop string) string {
	paths := []string{
		"/usr/share/applications/" + desktop,
		filepath.Join(os.Getenv("HOME"), ".local/share/applications/", desktop),
	}

	for _, p := range paths {
		if file, err := os.Open(p); err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "Exec=") {
					fields := strings.Fields(strings.TrimPrefix(line, "Exec="))
					return fields[0]
				}
			}
		}
	}

	return ""
}

func getBrowserVersion(_, path string) string {
	var out bytes.Buffer

	cmd := exec.Command(path, "--version")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return Unknown
	}

	return getDigits(out.String())
}
