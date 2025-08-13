# AtCoder workspace

AtCoder workspace.  This repository provides a preconfigured environment for participating in AtCoder contests using Go, C#, and other languages.

## Supported Languages

This workspace currently supports:

- Go (default)
- C#

## ⚙️ Setup

To get started instantly with a preconfigured environment using [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers), follow these steps:

### 1. Launch in DevContainer

1. Open this project in VS Code and run:

```bash
Cmd+Shift+P → Dev Containers: Reopen in Container
```

2. Restart VS Code

**Why**: The `acc` command will not be available in the terminal until VS Code is restarted after the container is built.

3. Run the VS Code task `[setup] After container setup`

### 2. Login to the tools

1. Login to [AtCoder](https://atcoder.jp/login) on your browser.  
2. Copy the REVEL_SESSION from developer tools -> Application -> Storage -> Cookies -> https://atcoder.jp -> REVEL_SESSION.
3. Run `aclogin`.  

```shell
$ aclogin
```

Paste copied REVEL_SESSION

4. Run VSCode task: `[setup] Update atcoder-cli template`
5. Complete

### 3. (Optional) Set up another language

With the above setup you can participate in contests using Go, but there’s limited support for a few other languages as well.  

Run the VS Code task `[setup] xx environment` (replace xx with the language name) to set up the environment for solving problems in that language.  

## 🚀 Usage

### Basic Workflow

1. **Create a new contest workspace:**

```bash
$ acc new abc123 --template go
```

2. **Write your solution:**
Edit the generated `main.go` file in each problem directory.

3. **Test your solution:**

Open Edited `main.go` file and run the test task: Run Test Cases

4. **Submit your solution:**

TODO: Not available yet. See issue: https://github.com/Tatamo/atcoder-cli/issues/68

```bash
$ acc submit
```

## 📁 Directory Structure

```plaintext
atcoder/
├── problems/                 # All contest problems live here
│   ├── {ContestID}/
│   │   ├── a/
├── scripts                   # utility scripts (call by vscode tasks)
├── settings                  # each tool settings
├── tools                     # utility tools
├── .gitignore
└── README.md
```

## License

This project is licensed under the MIT License.

It includes code from [gosagawa/atcoder](https://github.com/gosagawa/atcoder),
which is also licensed under the MIT License.
