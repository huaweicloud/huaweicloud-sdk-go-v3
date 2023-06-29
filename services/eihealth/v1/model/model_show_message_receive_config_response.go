package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageReceiveConfigResponse Response Object
type ShowMessageReceiveConfigResponse struct {

	// 接收范围（不接收消息(none)、仅接收自己操作的消息(mine)、接收全部消息(all)）
	Scope *string `json:"scope,omitempty"`

	// 资源类型
	ResourceTypes *[]string `json:"resource_types,omitempty"`

	Language       *LanguageEnum `json:"language,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowMessageReceiveConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageReceiveConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowMessageReceiveConfigResponse", string(data)}, " ")
}
