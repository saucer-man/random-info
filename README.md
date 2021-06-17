# random-info
个人信息生成器

# 目前支持

- 生成姓名
- 生成地址
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
	name, _ := info.RandomName()
	fmt.Printf("生成的姓名为：%s\n", name)

	phone, _ := info.RandomName()
	fmt.Printf("生成的手机号码为：%s\n", phone)

	addr, _ := info.RandomAddr()
	fmt.Printf("生成的地址为：%s\n", addr)

	idcard, _ := info.RandomIDcard("150701")
	fmt.Printf("生成的身份证号码为：%s\n", idcard)

	info, _ := info.RandomInfo()
	fmt.Printf("生成的个人信息为：%v\n", info)
}
```