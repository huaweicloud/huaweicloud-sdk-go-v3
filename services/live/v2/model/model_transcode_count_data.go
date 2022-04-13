package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeCountData struct {
	// 每个采样时间中的转码任务数信息。

	SpecList *[]TranscodeSpecCount `json:"spec_list,omitempty"`
	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ 。

	Time *string `json:"time,omitempty"`
}

func (o TranscodeCountData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeCountData struct{}"
	}

	return strings.Join([]string{"TranscodeCountData", string(data)}, " ")
}
