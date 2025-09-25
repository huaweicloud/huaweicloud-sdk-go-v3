package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDecoyPortAutoBindResponse Response Object
type ShowDecoyPortAutoBindResponse struct {

	// 是否自动绑定
	AutoBind *bool `json:"auto_bind,omitempty"`

	// 默认绑定的策略id
	WindowsPolicy *string `json:"windows_policy,omitempty"`

	// 默认绑定的策略id
	LinuxPolicy    *string `json:"linux_policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDecoyPortAutoBindResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDecoyPortAutoBindResponse struct{}"
	}

	return strings.Join([]string{"ShowDecoyPortAutoBindResponse", string(data)}, " ")
}
