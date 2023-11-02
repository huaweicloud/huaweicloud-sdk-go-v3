package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExpireAutoScanConfigRequest Request Object
type UpdateExpireAutoScanConfigRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateAutoScanConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateExpireAutoScanConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExpireAutoScanConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateExpireAutoScanConfigRequest", string(data)}, " ")
}
