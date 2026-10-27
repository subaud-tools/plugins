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
		{
			Name:        "muster",
			DisplayName: "muster",
			Bin:         "muster",
			ServerName:  "muster-channel",
			Description: "Mail on the muster bus starts a turn in the session it was sent to.",
			Author:      "Court Schuett",
			AuthorURL:   "https://muster.tools",
			Homepage:    "https://muster.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "BUSL-1.1",
			Keywords:    []string{"agents", "coordination", "tmux", "messaging"},
		},
		{
			Name:        "galley",
			DisplayName: "galley",
			Bin:         "galley",
			ServerName:  "galley",
			Description: "Review documents in the browser; rounds land in the session that wrote them.",
			Author:      "Court Schuett",
			AuthorURL:   "https://galley.tools",
			Homepage:    "https://galley.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "BUSL-1.1",
			Keywords:    []string{"review", "markdown", "writing", "editing"},
		},
		{
			Name:        "cull",
			DisplayName: "cull",
			Bin:         "cull",
			ServerName:  "cull",
			Description: "Judge your tests with Jev; Court's answers from cull's review page land in the session that asked.",
			Author:      "Court Schuett",
			AuthorURL:   "https://tackle.tools",
			Homepage:    "https://tackle.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "MIT",
			Keywords:    []string{"tests", "testing", "review"},
		},
		{
			Name:        "imprint",
			DisplayName: "imprint",
			Bin:         "imprint",
			ServerName:  "imprint",
			Description: "npm publishing, trusted publishers and deprecation, one step at a time; each step's progress lands in the session that started it.",
			Author:      "Court Schuett",
			AuthorURL:   "https://tackle.tools",
			Homepage:    "https://tackle.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "MIT",
			Keywords:    []string{"npm", "publishing", "trusted-publishing"},
		},
		{
			Name:        "sift",
			DisplayName: "sift",
			Bin:         "sift",
			ServerName:  "sift",
			Description: "Audit agent instruction files; the user's decisions from sift's review page land in the session that asked.",
			Author:      "Court Schuett",
			AuthorURL:   "https://tackle.tools",
			Homepage:    "https://tackle.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "MIT",
			Keywords:    []string{"instructions", "agents", "review"},
		},
		{
			Name:        "casebook",
			DisplayName: "casebook",
			Bin:         "casebook",
			ServerName:  "casebook",
			Description: "Decide the fate of every repo, pull request, issue, branch and worktree; the user's messages from casebook's page land in the session.",
			Author:      "Court Schuett",
			AuthorURL:   "https://casebook.tools",
			Homepage:    "https://casebook.tools",
			Repository:  "https://github.com/subaud-tools/plugins",
			License:     "MIT",
			Keywords:    []string{"github", "triage", "branches", "worktrees"},
		},
	} {
		t.Run(s.Name, func(t *testing.T) {
			claudeplugin.AssertCurrent(t, filepath.Join("plugins", s.Name), s)
		})
	}
}
