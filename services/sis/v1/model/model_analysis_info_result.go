package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnalysisInfoResult struct {
	// 角色类型, 目前仅支持 AGENT(座席), USER(用户)。

	Role *string `json:"role,omitempty"`
	// 情绪类型，目前支持NORMAL(正常)，ANGRY(愤怒)，UNKNOWN(未知)。 在识别配置中emotion为true时存在。

	Emotion *string `json:"emotion,omitempty"`
	// 语速信息，单位是\"每秒字数\"。 在识别配置中speed为true时存在。

	Speed *float32 `json:"speed,omitempty"`
}

func (o AnalysisInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisInfoResult struct{}"
	}

	return strings.Join([]string{"AnalysisInfoResult", string(data)}, " ")
}
