# FxxkGLfY

江西青年大学习记录生成器

# 用法

## 安装

```bash
go install github.com/Zxilly/FxxkGLfY@latest
```

## 运行

```bash
FxxkGlfY help
FxxkGlfy cfg
FxxkGlfY make
```

## 完整用法

```bash
NAME:
   FxxkGLfY - FxxkGLfY is a tool for generating a GLfY record. Current only for JiangXi Universities.


USAGE:
   FxxkGLfY [global options] command [command options] [arguments...]

COMMANDS:
   configure, cfg  Generating configuration
   make            Make a GLfY record with the configuration file or environment variables
   help, h         Shows a list of commands or help for one command


GLOBAL OPTIONS:
   --help, -h  show help (default: false)

NAME:
   FxxkGLfY make - Make a GLfY record with the configuration file or environment variables


USAGE:
   FxxkGLfY make [command options] [arguments...]

OPTIONS:
   --config value, -c value  Configuration file path to read, use "ENV" to read from environment variables (default: "token.json")

NAME:
   FxxkGLfY configure - Generating configuration

USAGE:
   FxxkGLfY configure [command options] [arguments...]

OPTIONS:
   --config value, -c value  Configuration file path to write, use "ENV" to show generated environment variables (default: "token.json")


```