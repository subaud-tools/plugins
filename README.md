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

## Add the marketplace

In Claude Code:

```
/plugin marketplace add subaud-tools/plugins
/plugin install diazo@subaud-tools
```

## Working on this repo

Every plugin is rendered by
[tools-common](https://github.com/subaud-tools/tools-common)'s `claudeplugin`
package, not written by hand. `plugins_test.go` holds each plugin's spec and
fails when a committed plugin differs from the render; regenerate with
`CLAUDEPLUGIN_UPDATE=1 go test ./...`. `just verify` runs that test and
Claude Code's own `claude plugin validate` on the marketplace and every
plugin, exactly as CI does.

## License

MIT, for the plugin files in this repo. Each tool a plugin runs has its own
license.
