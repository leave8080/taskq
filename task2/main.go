package main

import (
	"context"
	"fmt"
	"time"
)

//
//import (
//	"context"
//	"fmt"
//	"time"
//)
// 全局map
var jobmap = make(map[string]interface{})
//Job 任务
type Jobs struct {
	ID     string
	Status int
	Ctx    context.Context
	Cancel context.CancelFunc
}
func stopGetmi(id string) {
	//把任务停掉
	fmt.Println("stop jobs")
	jobss := jobmap[id]
	//interface 转 struct
	op, ok := jobss.(Jobs)
	fmt.Println(ok)
	// 调用砍头函数cancel
	defer op.Cancel()
}

func  main() {
		go test()
		time.Sleep(time.Second*10)
		go stopGetmi("sdads")
		time.Sleep(time.Second * 200)
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			// 接收到爹挂了的消息
			case <-ctx.Done():
				fmt.Println("儿子被砍头了。")
				// 退出任务
				return
			case dst <- n:
				n++
				time.Sleep(time.Second * 1)
			}
		}
	}()
	return dst
}
func test() {
	var jobs Jobs
	ctx, cancel := context.WithCancel(context.Background())
	// 造一个儿子
	intChan := gen(ctx)
	// 任务开始了
	fmt.Println("start job")
	// 重要的东西传进去
	jobs.Status = 1
	jobs.Cancel = cancel
	jobs.Ctx = ctx
	// 定义一个任务id，这个可以用uuid，或者随便整个别的
	jobs.ID = "sdads"
	jobmap["sdads"] = jobs
	// 阻塞任务，假装任务执行很久
	for n := range intChan {
		fmt.Println(n)
		if n == 1000 {
			break
		}
	}
}