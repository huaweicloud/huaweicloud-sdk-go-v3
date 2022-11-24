package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonthUsed struct {

	// 日期
	Date *string `json:"date,omitempty"`

	// 使用量
	Value *int32 `json:"value,omitempty"`
}

func (o MonthUsed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonthUsed struct{}"
	}

	return strings.Join([]string{"MonthUsed", string(data)}, " ")
}
