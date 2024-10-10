package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Point struct {

	// 时间戳
	Time *int32 `json:"time,omitempty"`

	// 请求总量
	Total *int32 `json:"total,omitempty"`

	// 攻击总量
	Attack *int32 `json:"attack,omitempty"`

	// web基础防护
	Basic *int32 `json:"basic,omitempty"`

	// 频率控制
	Cc *int32 `json:"cc,omitempty"`

	// 精准防护
	CustomCustom *int32 `json:"custom_custom,omitempty"`
}

func (o Point) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Point struct{}"
	}

	return strings.Join([]string{"Point", string(data)}, " ")
}
