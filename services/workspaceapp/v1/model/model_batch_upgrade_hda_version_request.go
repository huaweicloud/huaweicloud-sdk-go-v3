package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeHdaVersionRequest Request Object
type BatchUpgradeHdaVersionRequest struct {
	Body *BatchUpgradeHdaVersionReq `json:"body,omitempty"`
}

func (o BatchUpgradeHdaVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeHdaVersionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpgradeHdaVersionRequest", string(data)}, " ")
}
