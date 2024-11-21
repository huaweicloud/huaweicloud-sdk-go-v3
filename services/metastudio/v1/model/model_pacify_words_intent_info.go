package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PacifyWordsIntentInfo 安抚话术意图信息。
type PacifyWordsIntentInfo struct {

	// 意图名称
	Intent *string `json:"intent,omitempty"`

	// 意图中文描述
	DescCn *string `json:"desc_cn,omitempty"`

	// 意图英文描述
	DescEn *string `json:"desc_en,omitempty"`
}

func (o PacifyWordsIntentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PacifyWordsIntentInfo struct{}"
	}

	return strings.Join([]string{"PacifyWordsIntentInfo", string(data)}, " ")
}
