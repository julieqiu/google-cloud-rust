// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sidekick

import (
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
)

// generateClient takes some state and applies it to a template to create a client
// library.
func generateClient(data *TemplateData, root, outdir string) error {
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if outdir == "" {
				outdir, _ = os.Getwd()
			}
			dn := filepath.Join(outdir, strings.TrimPrefix(path, root))
			os.MkdirAll(dn, 0777) // Ignore errors
			return nil
		}
		if filepath.Ext(path) != ".mustache" {
			return nil
		}
		if strings.Count(d.Name(), ".") == 1 {
			// skipping partials
			return nil
		}
		s, err := mustache.RenderFile(path, data)
		if err != nil {
			return err
		}
		fn := filepath.Join(outdir, filepath.Dir(strings.TrimPrefix(path, root)), strings.TrimSuffix(d.Name(), ".mustache"))
		return os.WriteFile(fn, []byte(s), os.ModePerm)
	})
	if err != nil {
		slog.Error("error walking templates", "err", err.Error())
		return err
	}

	return nil
}
