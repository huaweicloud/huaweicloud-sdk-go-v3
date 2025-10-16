package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeConNumData struct {

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time *string `json:"time,omitempty"`

	// 转码路数。
	Value *int64 `json:"value,omitempty"`
}

func (o TranscodeConNumData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeConNumData struct{}"
	}

	return strings.Join([]string{"TranscodeConNumData", string(data)}, " ")
}
