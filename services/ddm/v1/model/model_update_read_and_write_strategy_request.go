package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateReadAndWriteStrategyRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *ModifyReadAndWriteStrategyReq `json:"body,omitempty"`
}

func (o UpdateReadAndWriteStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReadAndWriteStrategyRequest struct{}"
	}

	return strings.Join([]string{"UpdateReadAndWriteStrategyRequest", string(data)}, " ")
}
