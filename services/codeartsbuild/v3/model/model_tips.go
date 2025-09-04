package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tips 资讯
type Tips struct {

	// 状态
	Status *int32 `json:"status,omitempty"`

	// 下一个活动日
	NextAction *int32 `json:"next_action,omitempty"`

	// 下一个活动日剩余时间
	NextActionRemainDay *int32 `json:"next_action_remain_day,omitempty"`

	// 名称
	NextActionUrl *string `json:"next_action_url,omitempty"`
}

func (o Tips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tips struct{}"
	}

	return strings.Join([]string{"Tips", string(data)}, " ")
}
