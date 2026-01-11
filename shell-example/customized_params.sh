#!/bin/sh

# 定义脚本名称（用于错误提示）
SCRIPT_NAME=$(basename "$0")

# 定义默认参数值
INPUT_FILE=""
VERBOSE=false
OUTPUT_DIR="./output"

# 1. 定义支持的短选项和长选项
# 短选项：h(帮助)、f:(文件，需参数)、v(详细模式)、o:(输出目录，需参数)
# 长选项：help、file:、verbose、output:（与短选项一一对应）
SHORT_OPTS="hf:vo:"
LONG_OPTS="help,file:,verbose,output:"

# 2. 使用 getopt 规范化参数（关键步骤）
# -o 指定短选项，--long 指定长选项，-- 分隔选项和非选项参数
# -- "$@" 传递原始参数，避免特殊字符解析错误
PARSED_OPTS=$(getopt -o "$SHORT_OPTS" --long "$LONG_OPTS" --name "$SCRIPT_NAME" -- "$@")
if [ $? -ne 0 ]; then
    # getopt 解析失败（如无效选项），退出脚本
    echo "错误：参数解析失败，请使用 $SCRIPT_NAME --help 查看用法" >&2
    exit 1
fi

# 3. 替换原始参数为规范化后的参数
eval set -- "$PARSED_OPTS"

# 4. 循环解析参数
while true; do
    case "$1" in
        # 帮助选项
        -h|--help)
            echo "用法：$SCRIPT_NAME [选项] [非选项参数...]"
            echo "选项说明："
            echo "  -h, --help        显示帮助信息"
            echo "  -f, --file FILE   指定输入文件（必填）"
            echo "  -v, --verbose     开启详细输出模式"
            echo "  -o, --output DIR  指定输出目录（默认：./output）"
            exit 0
            ;;
        # 文件选项（带参数）
        -f|--file)
            INPUT_FILE="$2"
            shift 2  # 跳过选项和参数
            ;;
        # 详细模式
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        # 输出目录（带参数）
        -o|--output)
            OUTPUT_DIR="$2"
            shift 2
            ;;
        # 结束符（-- 之后的是非选项参数）
        --)
            shift
            break
            ;;
        # 未知选项（理论上 getopt 已过滤，兜底）
        *)
            echo "错误：未知选项 $1，请使用 --help 查看用法" >&2
            exit 1
            ;;
    esac
done

# 5. 参数校验（必填项检查）
if [ -z "$INPUT_FILE" ]; then
    echo "错误：必须通过 -f/--file 指定输入文件" >&2
    exit 1
fi

# 6. 非选项参数处理（如脚本需要额外的位置参数）
NON_OPTION_ARGS="$@"
if [ -n "$NON_OPTION_ARGS" ]; then
    echo "提示：非选项参数为：$NON_OPTION_ARGS"
fi

# ===================== 业务逻辑示例 =====================
echo "===== 脚本执行参数 ====="
echo "输入文件：$INPUT_FILE"
echo "详细模式：$VERBOSE"
echo "输出目录：$OUTPUT_DIR"

# 模拟业务逻辑
if [ "$VERBOSE" = true ]; then
    echo "详细模式：检查输入文件是否存在..."
    if [ -f "$INPUT_FILE" ]; then
        echo "成功：输入文件 $INPUT_FILE 存在"
    else
        echo "错误：输入文件 $INPUT_FILE 不存在" >&2
        exit 1
    fi
fi


