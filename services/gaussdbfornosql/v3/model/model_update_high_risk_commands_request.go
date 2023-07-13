package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHighRiskCommandsRequest Request Object
type UpdateHighRiskCommandsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RenameHighRiskCommandsRequest `json:"body,omitempty"`
}

func (o UpdateHighRiskCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHighRiskCommandsRequest struct{}"
	}

	return strings.Join([]string{"UpdateHighRiskCommandsRequest", string(data)}, " ")
}
