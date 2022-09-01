package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RealtimeScaleDimensionValue struct {

	// 维度值，如查询维度为region，则此处取值可能为GD
	Dimension *string `json:"dimension,omitempty" xml:"dimension"`

	// 在线观众数
	OnlineUsers *int64 `json:"online_users,omitempty" xml:"online_users"`
}

func (o RealtimeScaleDimensionValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealtimeScaleDimensionValue struct{}"
	}

	return strings.Join([]string{"RealtimeScaleDimensionValue", string(data)}, " ")
}
