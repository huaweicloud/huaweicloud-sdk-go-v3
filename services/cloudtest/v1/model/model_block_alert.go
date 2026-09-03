package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BlockAlert struct {
	AlertTemplate *AlertTemplate `json:"alert_template,omitempty"`

	// 阻塞告警开启 0关闭 1开启
	Enable *string `json:"enable,omitempty"`

	// 等待队列大于多少个开始阻塞
	WaitingCount *int32 `json:"waitingCount,omitempty"`
}

func (o BlockAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockAlert struct{}"
	}

	return strings.Join([]string{"BlockAlert", string(data)}, " ")
}
