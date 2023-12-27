package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSslSwitchRequest Request Object
type UpdateSslSwitchRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateSslSwitchRequestBody `json:"body,omitempty"`
}

func (o UpdateSslSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSslSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateSslSwitchRequest", string(data)}, " ")
}
