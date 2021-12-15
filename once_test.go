package zonce

import (
	"errors"
	"testing"
	"time"
)

var onces Once

func Test_once(t *testing.T)  {
	for{
		time.Sleep(time.Second)
		err := onces.Do(func() error {
			t.Log("执行一次")
			return errors.New("测试错误")
			//return errors.New("错误")
		})
		t.Log(err)

	}
}
