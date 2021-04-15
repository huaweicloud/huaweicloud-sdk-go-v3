package model

import (
	"encoding/json"

	"strings"
)

//
type ImageDetectionResultDetail struct {
	// 涉政敏感人物检测结果。

	Politics *[]ImageDetectionResultDetailPolitics `json:"politics,omitempty"`
	// 涉黄检测结果。

	Porn *[]ImageDetectionResultDetailPorn `json:"porn,omitempty"`
	// 涉政、暴恐检测结果。

	Terrorism *[]ImageDetectionResultDetailPorn `json:"terrorism,omitempty"`
	// 广告检测结果。

	Ad *[]ImageDetectionResultDetailPorn `json:"ad,omitempty"`
}

func (o ImageDetectionResultDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetail", string(data)}, " ")
}
