package genurl

import (
	"embed"
	"strings"
)

// 嵌入包含XML模板和其他资源的数据目录
//
//go:embed data
var DataDir embed.FS

// CleanName 清理产品名称，只保留英文字母并转为小写
//
// 此函数移除所有非英文字母的字符（如数字、标点、空格等），
// 并将结果转换为小写。适用于规范化产品名称。
//
// 参数:
//   - origin: 原始产品名称字符串
//
// 返回值:
//   - string: 清理后的小写字母字符串
//
// 示例:
//
//	name := CleanName("Microsoft.MinecraftUWP")  // 返回 "microsoftminecraftuwp"
func CleanName(origin string) string {
	var result strings.Builder

	for _, chr := range origin {
		if (64 < chr && chr < 91) || (96 < chr && chr < 123) {
			result.WriteRune(chr)
		}
	}

	return strings.ToLower(result.String())
}
