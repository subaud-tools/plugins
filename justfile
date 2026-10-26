# `just verify` is exactly what CI's `check` job runs.
set shell := ["bash", "-euo", "pipefail", "-c"]

default: verify

verify: test validate

# Each plugin is what tools-common renders for it (plugins_test.go).
# Regenerate with `CLAUDEPLUGIN_UPDATE=1 go test ./...`.
test:
    test -z "$(gofmt -l .)" || { gofmt -l .; echo "gofmt needed"; exit 1; }
    go vet ./...
    go test ./...

# The marketplace and each plugin pass Claude Code's own validator. Pinned, and
# run through npx, so CI's runner needs nothing installed and no login. Not
# --strict: the manifests carry no version on purpose (Claude Code then
# versions each plugin by commit), and the validator warns about that.
claude := "npx -y @anthropic-ai/claude-code@2.1.309"

validate:
    {{claude}} plugin validate .
    for p in plugins/*/; do {{claude}} plugin validate "$p"; done
