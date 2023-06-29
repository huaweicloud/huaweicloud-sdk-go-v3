package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMessageReceiveConfigReq 用户邮件接收配置请求体
type SetMessageReceiveConfigReq struct {

	// 接收范围（不接收消息(none)、仅接收自己操作的消息(mine)、接收全部消息(all)）
	Scope string `json:"scope"`

	// 资源类型
	ResourceTypes *[]string `json:"resource_types,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`
}

func (o SetMessageReceiveConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMessageReceiveConfigReq struct{}"
	}

	return strings.Join([]string{"SetMessageReceiveConfigReq", string(data)}, " ")
}
