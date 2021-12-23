package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpCode struct {
	// 状态码

	Code *int32 `json:"code,omitempty"`
	// 状态码出现次数

	Count *int32 `json:"count,omitempty"`
	// 状态码在对应时间点中的占比，保留4位小数。

	Proportion *float64 `json:"proportion,omitempty"`
}

func (o HttpCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpCode struct{}"
	}

	return strings.Join([]string{"HttpCode", string(data)}, " ")
}
