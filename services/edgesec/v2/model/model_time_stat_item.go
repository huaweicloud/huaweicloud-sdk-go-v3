package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeStatItem 时间统计模型
type TimeStatItem struct {

	// 时间戳，按时间粒度向左对齐。例如时间粒度为5分钟级，统计[19:10,19:15)，则该时间戳为19:10时刻时间
	Key int64 `json:"key"`

	// 攻击请求数
	Value int64 `json:"value"`
}

func (o TimeStatItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeStatItem struct{}"
	}

	return strings.Join([]string{"TimeStatItem", string(data)}, " ")
}
