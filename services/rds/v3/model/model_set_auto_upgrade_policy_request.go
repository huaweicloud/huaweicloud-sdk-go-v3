package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoUpgradePolicyRequest Request Object
type SetAutoUpgradePolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CustomerModifyAutoUpgradePolicyReq `json:"body,omitempty"`
}

func (o SetAutoUpgradePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoUpgradePolicyRequest struct{}"
	}

	return strings.Join([]string{"SetAutoUpgradePolicyRequest", string(data)}, " ")
}
