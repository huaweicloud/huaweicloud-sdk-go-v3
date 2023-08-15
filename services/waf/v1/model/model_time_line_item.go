package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeLineItem struct {

	// 时间点
	Time *int64 `json:"time,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`
}

func (o TimeLineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeLineItem struct{}"
	}

	return strings.Join([]string{"TimeLineItem", string(data)}, " ")
}
