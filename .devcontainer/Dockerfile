FROM mcr.microsoft.com/devcontainers/go:latest

# 必要なパッケージをインストール
RUN apt-get update && apt-get install -y \
    curl \
    nodejs \
    npm \
    python3 \
    python3-pip \
    python3-venv \
    python3-full \
    pipx \
    git \
    && rm -rf /var/lib/apt/lists/*

# Node.jsの最新LTS版にアップデート
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
    && apt-get install -y nodejs

# pipxのパス設定
ENV PATH="/root/.local/bin:$PATH"

USER vscode

# vscoderユーザーでpipxのパス設定
ENV PATH="/home/vscode/.local/bin:$PATH"
