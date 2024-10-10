package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceSpecRequest Request Object
type UpgradeInstanceSpecRequest struct {
	Body *UpgradeInstanceSpecV2RequestBody `json:"body,omitempty"`
}

func (o UpgradeInstanceSpecRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceSpecRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceSpecRequest", string(data)}, " ")
}
