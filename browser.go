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
	"fmt"
)

const (
	Unknown = "Unknown"
)

const (
	Chrome   = "Chrome"
	Chromium = "Chromium"
	Edge     = "Edge"
	Firefox  = "Firefox"
	Safari   = "Safari"
)

type Browser struct {
	Name    string
	Path    string
	Version string
}

func (b Browser) String() string {
	return fmt.Sprintf("Name: %s\nPath: %s\nVersion: %s", b.Name, b.Path, b.Version)
}

// OpenFile open the given html file in OS default browser window
func OpenFile(path string) error {
	return openFile(path)
}

// OpenFileByBrowser open the given browser window and then open the html file in the browser window
func OpenFileByBrowser(browser Browser, path string) error {
	return openFileByBrowser(browser, path)
}

// OpenURL open a browser window and redirect to the given url.
func OpenURL(url string) error {
	return openURL(url)
}

// OpenURLByBrowser open the given browser window and redirect to the given url.
func OpenURLByBrowser(browser Browser, url string) error {
	return openURLByBrowser(browser, url)
}

// DefaultBrowser detect the OS default browser
func DefaultBrowser() Browser {
	return detectDefaultBrowser()
}

// InstalledBrowsers only detected the default browser paths like Chrome, Chromium, Edge, Firefox and Safari
// the first one is the default browser in current OS.
func InstalledBrowsers() []Browser {
	return detectInstalledBrowsers()
}
