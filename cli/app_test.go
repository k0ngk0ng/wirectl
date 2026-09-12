package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestRouting(t *testing.T) {
	var got []string
	app := App{Name: "test", Commands: map[string]Command{"download": {Run: func(ctx context.Context, args []string) error { got = args; return nil }}}}
	if err := app.Run(context.Background(), []string{"download", "magnet:?a=b&c=d"}); err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0] != "magnet:?a=b&c=d" {
		t.Fatal(got)
	}
}
func TestHelpAndUnknown(t *testing.T) {
	var out bytes.Buffer
	app := App{Name: "wirectl", Out: &out, PluginPrefix: "not-installed-"}
	if err := app.Run(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), "PATH") {
		t.Fatal(out.String())
	}
	if err := app.Run(context.Background(), []string{"../bad"}); err == nil {
		t.Fatal("invalid plugin name accepted")
	}
}
