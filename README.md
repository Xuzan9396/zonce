### zonce
##### 官方的sync.Once 不管有没有错误，始终执行一次，但是会出现一种问题有时候经过网络io操作的，可能这个时候出现错误，例如数据库,或者其他中间件刚刚好网络异常，或者网络因素导致了初始化失败，这个时候代码就会有问题
##### 自定义once包，解决初始化错误的情况，返回了err,如果有err，下次执行还会访问once里面内容


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
