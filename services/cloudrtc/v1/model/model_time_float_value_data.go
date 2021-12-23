package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeFloatValueData struct {
	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ

	Time *string `json:"time,omitempty"`
	// 当前时间返回参数取值

	Value *float32 `json:"value,omitempty"`
}

func (o TimeFloatValueData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeFloatValueData struct{}"
	}

	return strings.Join([]string{"TimeFloatValueData", string(data)}, " ")
}
