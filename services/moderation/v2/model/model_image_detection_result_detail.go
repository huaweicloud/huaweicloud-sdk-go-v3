package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ImageDetectionResultDetail struct {
	// 涉政敏感人物检测结果。

	Politics *[]ImageDetectionResultDetailPolitics `json:"politics,omitempty"`
	// 涉黄检测结果。

	Porn *[]ImageDetectionResultSimpleDetail `json:"porn,omitempty"`
	// 涉政、暴恐检测结果。

	Terrorism *[]ImageDetectionResultSimpleDetail `json:"terrorism,omitempty"`
	// 广告检测结果。

	Ad *[]ImageDetectionResultAdDetail `json:"ad,omitempty"`
}

func (o ImageDetectionResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetail", string(data)}, " ")
}
