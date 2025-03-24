package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟一个耗时的任务
func longRunningTask(ctx context.Context) (string, error) {
	// select {
	// case <-time.After(1 * time.Second): // 模拟任务耗时 5 秒
	// 	return "Task completed!", nil
	// case <-ctx.Done(): // 如果 context 被取消或超时
	// 	return "", ctx.Err() // 返回 context 的错误
	// }
}

func main() {
	// 创建一个带有超时的 context，超时时间为 3 秒
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 确保 context 被取消，释放资源

	// 启动耗时任务
	result, err := longRunningTask(ctx)
	if err != nil {
		fmt.Println("Error:", err) // 输出超时或取消的错误
		return
	}

	fmt.Println(result) // 输出任务结果
}
