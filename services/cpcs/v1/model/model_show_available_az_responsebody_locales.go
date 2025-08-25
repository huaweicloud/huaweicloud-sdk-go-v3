package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableAzResponsebodyLocales 区域名称
type ShowAvailableAzResponsebodyLocales struct {

	// 英文
	EnUs *string `json:"en-us,omitempty"`

	// 中文
	ZhCn *string `json:"zh-cn,omitempty"`
}

func (o ShowAvailableAzResponsebodyLocales) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableAzResponsebodyLocales struct{}"
	}

	return strings.Join([]string{"ShowAvailableAzResponsebodyLocales", string(data)}, " ")
}
