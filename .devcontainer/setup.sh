#!/bin/bash

# NPMのグローバルディレクトリをユーザーローカルに設定
mkdir -p ~/.npm-global
npm config set prefix '~/.npm-global'

# AtCoder周辺ツールをインストール
npm install -g atcoder-cli
pipx install online-judge-tools
pipx install aclogin

# PATHの設定（.bashrcに追加）
echo 'export PATH="$HOME/.npm-global/bin:$HOME/.local/bin:$PATH"' >> ~/.bashrc

# 現在のセッション用にPATHを設定
export PATH="$HOME/.npm-global/bin:$HOME/.local/bin:$PATH"

# pipxのパス確認
pipx ensurepath

# atcoder-cli settings
acc config default-contest-dirname-format /workspaces/atcoder/problems/{ContestID}
cp -r /workspaces/atcoder/settings/acc/* `acc config-dir`
acc config default-template go

# GNU time Not available とか出るのでインストールしておく
sudo apt update && sudo apt install time -y
