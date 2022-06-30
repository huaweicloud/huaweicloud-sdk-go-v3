package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 攻击事件统计结果
type CountItem struct {

	// 类型
	Key *string `json:"key,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`

	// 域名
	Host *string `json:"host,omitempty"`
}

func (o CountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountItem struct{}"
	}

	return strings.Join([]string{"CountItem", string(data)}, " ")
}
