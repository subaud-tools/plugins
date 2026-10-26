package plugins

import (
	"path/filepath"
	"testing"

	"github.com/subaud-tools/tools-common/claudeplugin"
)

// Each plugin is rendered by tools-common, not written by hand: its manifest,
// the .mcp.json that runs `<tool> channel` for the tools, and the mod that
// runs `<tool> watch` and starts a turn with each event. Regenerate with
// CLAUDEPLUGIN_UPDATE=1 go test ./...
func TestPluginsAreCurrent(t *testing.T) {
	for _, s := range []claudeplugin.Spec{
		{
			Name:        "diazo",
			DisplayName: "diazo",
			Bin:         "diazo",
			ServerName:  "diazo",
			Description: "Plan with an agent on a page in the browser; what you decide there lands in the session that opened the folder.",
			Author:      "Court Schuett",
			Homepage:    "https://diazo.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "MIT",
			Keywords:    []string{"planning", "design", "spec"},
		},
	} {
		t.Run(s.Name, func(t *testing.T) {
			claudeplugin.AssertCurrent(t, filepath.Join("plugins", s.Name), s)
		})
	}
}
