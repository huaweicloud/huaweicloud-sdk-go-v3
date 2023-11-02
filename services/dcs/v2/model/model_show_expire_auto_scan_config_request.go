package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExpireAutoScanConfigRequest Request Object
type ShowExpireAutoScanConfigRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowExpireAutoScanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExpireAutoScanConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowExpireAutoScanConfigRequest", string(data)}, " ")
}
