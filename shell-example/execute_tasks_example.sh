#!/bin/sh

# 定义脚本名称（用于错误提示）
SCRIPT_NAME=$(basename "$0")

TASKS=''
OUTPUT_DIR=""

SHORT_OPTS="ht:o:"
LONG_OPTS='help,tasks:,output:'

PARSED_OPTS=$(getopt -o "$SHORT_OPTS" --long "$LONG_OPTS" --name "$SCRIPT_NAME" -- "$@")

if [ "$?" -ne 0 ]; then
  # getopt 解析失败（如无效选项），退出脚本
  echo "错误：参数解析失败，请使用 $SCRIPT_NAME --help 查看用法" >&2
  exit 1
fi

eval set -- "$PARSED_OPTS"
echo PARSED_OPTS: $PARSED_OPTS

while true; do
  case "$1" in
    --help)
      echo "用法：$SCRIPT_NAME [选项] [非选项参数...]"
      echo "选项说明："
      echo "  -h, --help        显示帮助信息"
      echo "  -t --tasks FILE   指定输入文件（必填）"
      echo "  -o, --output DIR  指定输出目录（默认：./output）"
      exit 0
      ;;
    -t|--tasks)
      TASKS=$2
      shift 2;
      ;;
    -o|--output)
      OUTPUT_DIR=$2;
      shift 2;
      ;;
    --)
      shift 2;
      break;
      ;;
      *)
      echo "错误：未知选项 $1，请使用 --help 查看用法" >&2
      exit 1
      ;;
  esac
done

echo params;
echo TASKS: $TASKS
echo OUTPUT_DIR:$OUTPUT_DIR

IFS=',' read -ra tasks <<< "$TASKS";

index=0;

while [ $index -lt "${#tasks[@]}" ]; do
  echo "index $index value: ${tasks[index]}";
  index=$(expr $index + 1)
done;



