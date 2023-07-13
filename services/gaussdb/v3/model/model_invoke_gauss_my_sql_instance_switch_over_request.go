package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeGaussMySqlInstanceSwitchOverRequest Request Object
type InvokeGaussMySqlInstanceSwitchOverRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *TaurusSwitchoverRequest `json:"body,omitempty"`
}

func (o InvokeGaussMySqlInstanceSwitchOverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeGaussMySqlInstanceSwitchOverRequest struct{}"
	}

	return strings.Join([]string{"InvokeGaussMySqlInstanceSwitchOverRequest", string(data)}, " ")
}
