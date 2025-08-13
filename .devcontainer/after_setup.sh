#!/usr/bin/env bash
set -euo pipefail

# Get the workspace root directory (default: /workspaces/atcoder)
ROOT="${1:-/workspaces/atcoder}"

# Configure the contest directory format based on the workspace
acc config default-contest-dirname-format "$ROOT/problems/{ContestID}"

# Copy configuration files (including hidden files)
cfgdir="$(acc config-dir)"
mkdir -p "$cfgdir"
cp -r "$ROOT/settings/acc/." "$cfgdir/"

# Set the default template
acc config default-template go