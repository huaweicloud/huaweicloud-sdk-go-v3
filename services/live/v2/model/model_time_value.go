package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeValue struct {
	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ

	Time *string `json:"time,omitempty"`
	// 当前时间返回指定指标的值

	Value *int64 `json:"value,omitempty"`
}

func (o TimeValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeValue struct{}"
	}

	return strings.Join([]string{"TimeValue", string(data)}, " ")
}
