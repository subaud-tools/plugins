# galley plugin

The plugin carries three things: the `.mcp.json` entry pointing at `galley channel` on PATH, which serves galley's tools; the mod in `hooks/`, which runs `galley watch` and starts a turn with each event; and the manifest that lets them install in one step. All three are rendered by tools-common (`claudeplugin`), and `TestPluginsAreCurrent` (`plugins_test.go`) fails when the committed copy differs.

The skill that teaches an agent the review loop, `reviewing-a-document`, is not in the plugin. It is embedded in the galley binary: `galley skills install` writes it for each agent harness it finds (Claude Code, Codex, pi), and `galley channel` keeps each copy current each time it starts.

**The binary is the product and the licence boundary.** The plugin ships no galley code and grants no entitlement. `galley` must already be on PATH.

## Installing

The plugin is in the subaud-tools marketplace, `subaud-tools/plugins`.

```sh
claude plugin marketplace add subaud-tools/plugins
claude plugin install galley@subaud-tools
galley skills install
```

The commands work for anyone with `galley` on PATH. A session started after the install hears galley with plain `claude`. In a session that is already open, run `/reload-plugins` to load it.

Turn on auto-update for the `subaud-tools` marketplace so the plugin stays current: in Claude Code, run `/plugin`, open the **Marketplaces** tab, select `subaud-tools` and choose **Enable auto-update**.

## Versions

A version sweep will find one number and one place without one, and they are not the same thing. Nothing here should ever be "kept in sync".

| Where | What it is | Moves when |
|---|---|---|
| galley's `VERSION` (in the galley repo) | galley the binary | a galley release is cut |
| `plugins/galley/.claude-plugin/plugin.json` | **no version at all**: Claude Code versions the plugin by the commit it was installed from | any commit that changes the plugin |

The plugin does not track the binary. It ships a manifest, an MCP entry and a mod, not galley code, and its manifest carries no version on purpose, so nobody has to bump one.

## Opening a document

The channel opens documents for its own session through the `galley_open` tool. It starts `galley edit <doc> --owner <session>` detached, which opens the page in the reviewer's browser, waits for the editor's advert, and returns the URL. A document the session already has open returns the same URL and opens nothing. The editor is bound to the session's channel presence and stops itself a few seconds after the session ends. An agent must not run `galley edit` from a shell while the channel is present: an editor opened that way carries no owner, and the first channel whose scope covers it claims it. `galley edit` without `--owner` is the human's command. The skill says so, and so do the channel instructions.

## Without the plugin

In an agent without the plugin or channels.tools, `galley wait <doc>` is the same events over a blocking pull, and `galley agent-prompt <doc>` prints the protocol for an agent starting cold. Neither needs a plugin or Claude Code.
