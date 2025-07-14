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
	"testing"
)

func TestOpenFile(t *testing.T) {
	err := OpenFile("browser.html")
	if err != nil {
		t.Fatalf("Open File failed: %v", err)
	}
}

func TestOpenFileByBrowser(t *testing.T) {
	browsers := InstalledBrowsers()

	for _, browser := range browsers {
		err := openFileByBrowser(browser, "browser.html")
		if err != nil {
			t.Fatalf("Open File failed by browser[%v]: %v", browser, err)
		}
	}
}

func TestOpenURL(t *testing.T) {
	err := OpenURL("https://www.google.com")
	if err != nil {
		t.Fatalf("Open URL failed: %v", err)
	}
}

func TestOpenURLByBrowser(t *testing.T) {
	browsers := InstalledBrowsers()

	for _, browser := range browsers {
		err := openURLByBrowser(browser, "https://www.google.com")
		if err != nil {
			t.Fatalf("Open URL failed by browser[%v]: %v", browser, err)
		}
	}
}

func TestDefaultBrowser(t *testing.T) {
	defaultBrowser := DefaultBrowser()

	t.Logf("DefaultBrowser:\n%v", defaultBrowser)
}

func TestInstalledBrowsers(t *testing.T) {
	installedBrowsers := InstalledBrowsers()

	t.Log("InstalledBrowsers:")
	for _, browser := range installedBrowsers {
		t.Logf("\n%v", browser)
	}
}
