package stores

import (
	"go/parser"
	"go/token"

	"strconv"

	"sort"

	"github.com/dave/flux"
	"github.com/dave/jsgo/playground/actions"
)

func NewScannerStore(app *App) *ScannerStore {
	s := &ScannerStore{
		app: app,
	}
	return s
}

type ScannerStore struct {
	app     *App
	imports []string
}

// Imports returns a snapshot of imports
func (s *ScannerStore) Imports() []string {
	var a []string
	for _, v := range s.imports {
		a = append(a, v)
	}
	return a
}

func (s *ScannerStore) Changed(compare []string) bool {
	if len(compare) != len(s.imports) {
		return true
	}
	for i := range compare {
		if s.imports[i] != compare[i] {
			return true
		}
	}
	return false
}

func (s *ScannerStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.UserChangedText:
		fset := token.NewFileSet()

		// ignore errors
		f, _ := parser.ParseFile(fset, s.app.Editor.Current(), action.Text, parser.ImportsOnly)

		var imports []string
		for _, v := range f.Imports {
			// ignore errors
			unquoted, _ := strconv.Unquote(v.Path.Value)
			imports = append(imports, unquoted)
		}
		sort.Strings(imports)

		if s.Changed(imports) {
			s.imports = imports
			s.app.Debug("Imports", s.imports)
			s.app.Dispatch(&actions.ImportsChanged{})
			payload.Notify()
		}
	}
	return true
}
