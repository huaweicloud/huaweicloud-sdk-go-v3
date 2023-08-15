package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableQueueInfo struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	Status *string `json:"status,omitempty"`

	// 队列名称。
	Name *string `json:"name,omitempty"`

	// uuid。
	Uuid *string `json:"uuid,omitempty"`

	// 状态为失败时的详细报错信息。
	ErrMsg *string `json:"err_msg,omitempty"`

	// 作业更新时间, 毫秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AvailableQueueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableQueueInfo struct{}"
	}

	return strings.Join([]string{"AvailableQueueInfo", string(data)}, " ")
}
