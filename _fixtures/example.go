package foo

import "fmt"

const (
	// 这是comment单行
	foo = "这是一个Go的提示信息" // 测试after注释
	/* 测试multi-line字符串 */
	bar = `这是Go长字符串
	这是string第二行`
)

func main() {
	fmt.Println("准备打印Foo的结果:", foo)
}
