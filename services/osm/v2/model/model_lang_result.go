package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LangResult struct {

	// 语言
	Language *string `json:"language,omitempty"`

	// 是否英语
	IsEn *bool `json:"is_en,omitempty"`
}

func (o LangResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LangResult struct{}"
	}

	return strings.Join([]string{"LangResult", string(data)}, " ")
}
