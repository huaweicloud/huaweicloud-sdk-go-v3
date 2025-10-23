package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestartPolicyRequest Request Object
type ShowRestartPolicyRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowRestartPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestartPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowRestartPolicyRequest", string(data)}, " ")
}
