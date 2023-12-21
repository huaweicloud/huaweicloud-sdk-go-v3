package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceRequest Request Object
type UpgradeInstanceRequest struct {
	Body *UpgradeCbhRequestBody `json:"body,omitempty"`
}

func (o UpgradeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceRequest", string(data)}, " ")
}
