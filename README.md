# subaud-tools plugins

The Claude Code plugin marketplace for the subaud .tools family. Each plugin
connects Claude Code to one family tool: it registers the tool's MCP server
(the tool's `channel` command) and a mod that runs the tool's `watch` command,
so what you do on the tool's page in the browser arrives in your Claude Code
session as a turn.

The plugins hold no product code: each one only runs a tool you install
yourself. Install the tool from its site first; a plugin without its tool on
your `PATH` does nothing.

| Plugin | What it does | Install the tool from |
| --- | --- | --- |
| `diazo` | Plan and design with an agent on a page in the browser. | [diazo.tools](https://diazo.tools) |
| `muster` | Mail on the muster bus starts a turn in the session it was sent to. | [muster.tools](https://muster.tools) |
| `galley` | Review a Markdown document in the browser while an agent revises it; its skill comes from `galley skills install`. | [galley.tools](https://galley.tools) |
| `casebook` | Decide the fate of every repo, pull request, issue, branch and worktree from casebook's page. | [casebook.tools](https://casebook.tools) |
| `cull` | Judge a project's tests with Jev; your answers from cull's review page land in the session. | [tackle.tools](https://tackle.tools) |
| `imprint` | npm publishing, trusted publishers and deprecation, one step at a time. | [tackle.tools](https://tackle.tools) |
| `sift` | Audit agent instruction files; your decisions from sift's review page land in the session. | [tackle.tools](https://tackle.tools) |

## Add the marketplace

In Claude Code:

```
/plugin marketplace add subaud-tools/plugins
/plugin install diazo@subaud-tools
```

Install any other plugin the same way, as `<plugin>@subaud-tools`.

## Working on this repo

Every plugin's manifest, `.mcp.json` and mod are rendered by
[tools-common](https://github.com/subaud-tools/tools-common)'s `claudeplugin`
package, not written by hand. No plugin carries a skill: a tool's skill is
embedded in its binary and installed with `<tool> skills install` (galley's
`reviewing-a-document`, diazo's `diazo`). `plugins_test.go` holds each plugin's spec and
fails when a committed plugin differs from the render; regenerate with
`CLAUDEPLUGIN_UPDATE=1 go test ./...`. `just verify` runs that test and
Claude Code's own `claude plugin validate` on the marketplace and every
plugin, exactly as CI does.

## License

MIT, for every file in this repo, every plugin included. Each tool a plugin
runs has its own license.
