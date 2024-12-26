package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProviderValueBody 运营商名称
type ProviderValueBody struct {

	// 英文名称。
	EnUs *string `json:"en_us,omitempty"`

	// 中文信息。
	ZhCn *string `json:"zh_cn,omitempty"`
}

func (o ProviderValueBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderValueBody struct{}"
	}

	return strings.Join([]string{"ProviderValueBody", string(data)}, " ")
}
