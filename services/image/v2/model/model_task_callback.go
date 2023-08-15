package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskCallback struct {

	// 回调url地址，用于通知客户任务运行结束。
	Url *string `json:"url,omitempty"`
}

func (o TaskCallback) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCallback struct{}"
	}

	return strings.Join([]string{"TaskCallback", string(data)}, " ")
}
