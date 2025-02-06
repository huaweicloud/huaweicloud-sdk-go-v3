package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoUpgradePolicyRequest Request Object
type ShowAutoUpgradePolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowAutoUpgradePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoUpgradePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoUpgradePolicyRequest", string(data)}, " ")
}
