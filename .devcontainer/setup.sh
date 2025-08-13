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

# GNU time Not available とか出るのでインストールしておく
sudo apt update && sudo apt install time -y
