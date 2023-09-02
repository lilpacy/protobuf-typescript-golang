package main

import (
	"fmt"

	"github.com/lilpacy/protobuf-typescript-golang/common"
)

func doSomething(status common.Status) {
	switch status {
	case common.Status_ACTIVE:
		fmt.Println("ステータスはアクティブです")
	case common.Status_INACTIVE:
		fmt.Println("ステータスは非アクティブです")
	case common.Status_UNKNOWN:
		fmt.Println("ステータスは不明です")
	default:
		panic("未知のステータス")
	}
}

func main() {
	doSomething(common.Status_ACTIVE)
}
