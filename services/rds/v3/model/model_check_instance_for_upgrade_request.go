package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckInstanceForUpgradeRequest Request Object
type CheckInstanceForUpgradeRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *RdsUpgradePrecheckV3Req `json:"body,omitempty"`
}

func (o CheckInstanceForUpgradeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckInstanceForUpgradeRequest struct{}"
	}

	return strings.Join([]string{"CheckInstanceForUpgradeRequest", string(data)}, " ")
}
