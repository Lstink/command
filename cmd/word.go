package cmd

import (
	"github.com/lstink/command/internel/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部打次转化为大写",
	"2：全部单词转化为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "Word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCameCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCameCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCameCaseToUnderscore:
			content = word.CameCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请先执行 help word 查看帮助文档")
		}

		log.Printf("输出结果：%s", content)
	},
}

const (
	ModeUpper                     = iota + 1 // 全部单词转化为大写
	ModeLower                                // 全部单词转化为小写
	ModeUnderscoreToUpperCameCase            // 下划线单词转为大写驼峰单词
	ModeUnderscoreToLowerCameCase            // 下划线单词转为小写驼峰单词
	ModeCameCaseToUnderscore                 // 驼峰单词转为下划线单词
)

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
