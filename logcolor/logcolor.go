package logcolor

import (
	"fmt"
	"runtime"
)

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textMagenta
	textCyan
	textWhite
)

// Success 成功输出样式
func Success(str string) string {
	return textColor(textGreen, "[SUCCESS]") + fmt.Sprintf(" %s\n", str)
}

// Warning 警告输出样式
func Warning(str string) string {
	return textColor(textYellow, "[WARNING]") + fmt.Sprintf(" %s\n", str)
}

// Error 错误输出样式
func Error(str string) string {
	return textColor(textRed, "[ERROR]") + fmt.Sprintf(" %s\n", str)
}

// Info 错误输出样式
func Info(str string) string {
	return textColor(textBlue, "[INFO]") + fmt.Sprintf(" %s\n", str)
}

// Black 颜色输出
func Black(str string) string {
	return textColor(textBlack, str)
}

// Red 颜色输出
func Red(str string) string {
	return textColor(textRed, str)
}

// Green 颜色输出
func Green(str string) string {
	return textColor(textGreen, str)
}

// Yellow 颜色输出
func Yellow(str string) string {
	return textColor(textYellow, str)
}

// Blue 颜色输出
func Blue(str string) string {
	return textColor(textBlue, str)
}

// Magenta 颜色输出
func Magenta(str string) string {
	return textColor(textMagenta, str)
}

// Cyan 颜色输出
func Cyan(str string) string {
	return textColor(textCyan, str)
}

// White 颜色输出
func White(str string) string {
	return textColor(textWhite, str)
}

// IsWindows 验证场景
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func textColor(color int, str string) string {
	if IsWindows() {
		return str
	}
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, str)
}
