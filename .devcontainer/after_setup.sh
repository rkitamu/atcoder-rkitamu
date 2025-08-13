#!/usr/bin/env bash
set -euo pipefail

# 引数で workspace ルートを受け取る（未指定なら /workspaces/atcoder を仮定）
ROOT="${1:-/workspaces/atcoder}"

# acc 設定：コンテストのディレクトリ形式をワークスペース基準に
acc config default-contest-dirname-format "$ROOT/problems/{ContestID}"

# 設定ファイルコピー（dotfileも含めて安全に）
cfgdir="$(acc config-dir)"
mkdir -p "$cfgdir"
# 末尾を「/.」にすると * と違って隠しファイルも含めて中身をコピーできる
cp -r "$ROOT/settings/acc/." "$cfgdir/"

# デフォルトテンプレート
acc config default-template go