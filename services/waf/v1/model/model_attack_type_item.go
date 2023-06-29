package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttackTypeItem 攻击事件分布统计
type AttackTypeItem struct {

	// 攻击事件类型
	Key *string `json:"key,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`
}

func (o AttackTypeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackTypeItem struct{}"
	}

	return strings.Join([]string{"AttackTypeItem", string(data)}, " ")
}
