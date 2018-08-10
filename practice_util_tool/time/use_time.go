package time

import (
	"time"
	"fmt"
)

func main() {
	usetime()
}

func usetime(){

	//获取当前的时间戳
	now := time.Now()
	secondNum := now.Unix()
	fmt.Println(secondNum)

	//时间戳转换成时间
	now1 := time.Unix(secondNum, 0)
	fmt.Println(now1.Format("2006-01-02 15:04:05"))

	//格式化当前的时间
	tshow := now.Format("2006-01-02 15:04:05")
	fmt.Println(tshow)

	//str转换成时间time
	t, _ := time.Parse("2006-01-02 15:04:05", "2018-08-04 08:37:18")
	fmt.Println("t(time format)", t)

	//时间的增删改查
	fmt.Println(t.Add(time.Second * 10))
	fmt.Println(t.Sub(now1).Hours())

	c := make(chan int)
	//after time
	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
	}
	//定时器
	atick := time.Tick(1 * time.Second)
	select {
	case <-atick:
		fmt.Println("tt")
	}
	/*
	for now := range atick {
		fmt.Printf("%v %s\n", now)
	}*/


}
