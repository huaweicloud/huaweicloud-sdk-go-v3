package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiApiKeyCreate struct {

	// AIAPIKey值，不指定具体值时，由后台自动生成随机字符串。 支持大小写英文字母、数字，以及+-_/=特殊字符，长度为8~128个字符。
	AiApiKey *string `json:"ai_api_key,omitempty"`

	// AIAPIKey的别名。支持大小写字母，数字，下划线，中划线，长度为1~100个字符。
	Alias *string `json:"alias,omitempty"`
}

func (o AiApiKeyCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiApiKeyCreate struct{}"
	}

	return strings.Join([]string{"AiApiKeyCreate", string(data)}, " ")
}
