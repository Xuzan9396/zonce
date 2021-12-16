# zonce
自定义once包，解决初始化错误的情况


```go
package main

import (
	"errors"
	"log"
)
import "github.com/Xuzan9396/zonce"

func main() {
	var onces zonce.Once
    // 第一次执行中出现了失败,模拟初始化失败了
	err := onces.Do(func() error {
		log.Println("执行一次1")
		return errors.New("模拟实例化中间出现错误")
	})
	log.Println(err)

	// 第二次成功了
	err = onces.Do(func() error {
		log.Println("执行一次2")
		return nil
	})
	log.Println(err)


	// 第三次不会重新实例化了
	err = onces.Do(func() error {
		log.Println("执行一次3")
		return nil
	})
	log.Println(err)

}

```