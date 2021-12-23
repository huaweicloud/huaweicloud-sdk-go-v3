package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeValueData struct {
	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ

	Time *string `json:"time,omitempty"`
	// 查询的流量指标值

	Value *int64 `json:"value,omitempty"`
}

func (o TimeValueData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeValueData struct{}"
	}

	return strings.Join([]string{"TimeValueData", string(data)}, " ")
}
