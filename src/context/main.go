package main

import (
	"fmt"
	"time"
)

// // 模拟一个耗时的任务
// func longRunningTask(ctx context.Context) (string, error) {
// 	// select {
// 	// case <-time.After(1 * time.Second): // 模拟任务耗时 5 秒
// 	// 	return "Task completed!", nil
// 	// case <-ctx.Done(): // 如果 context 被取消或超时
// 	// 	return "", ctx.Err() // 返回 context 的错误
// 	// }
// }

// func main() {
// 	// 创建一个带有超时的 context，超时时间为 3 秒
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel() // 确保 context 被取消，释放资源

// 	// 启动耗时任务
// 	result, err := longRunningTask(ctx)
// 	if err != nil {
// 		fmt.Println("Error:", err) // 输出超时或取消的错误
// 		return
// 	}

// 	fmt.Println(result) // 输出任务结果
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	// defer cancel() ctx.Err() 为nil
// 	cancel() // ctx.Err() 为 context canceled

// 	fmt.Println("ctx.Err(): ", ctx.Err())

// 	now := time.Now()
// 	d := now.Add(5 * time.Minute)
// 	nCtx, deadlineCancel := context.WithDeadline(context.Background(), d)
// 	defer deadlineCancel()
// 	fmt.Println("now: ", now.Unix(), " d: ", d.Unix())

// 	deadline, ok := nCtx.Deadline()
// 	if ok {
// 		fmt.Println("deadline: ", deadline.Unix())
// 	}
// }

func Done() <-chan struct{} {
	return nil
}

func main() {
	fmt.Println("test start")

	go func() {
		fmt.Println("in select")
		select {
		case <-Done():
			fmt.Println("nil chan")
		}
	}()

	time.Sleep(5 * time.Second)

	fmt.Println("test end")

}
