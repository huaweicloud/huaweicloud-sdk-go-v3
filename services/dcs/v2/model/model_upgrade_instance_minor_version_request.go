package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceMinorVersionRequest Request Object
type UpgradeInstanceMinorVersionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpgradeMinorVersionRequestBody `json:"body,omitempty"`
}

func (o UpgradeInstanceMinorVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceMinorVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceMinorVersionRequest", string(data)}, " ")
}
