# random-info
个人信息生成器

# 目前支持

- 生成姓名
- 生成地址

todo
- 手机号
- 身份证号码

## 如何使用

```
package main

import(
    "fmt"
    info "github.com/saucer-man/random-info"
)

func main() {
	fmt.Println("hello")
	name, err := info.RandomName()
	if err != nil {
		fmt.Println("生成随机姓名出错")
	}
	fmt.Printf("生成的姓名为：%s\n", name)

	addr, err := info.RandomAddr()
	if err != nil {
		fmt.Println("生成随机地址出错")
	}
	fmt.Printf("生成的地址为：%s\n", addr)
}
```