package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceGroupRequest Request Object
type UpdateInstanceGroupRequest struct {

	// 引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *BatchModifyGroupDescriptionReq `json:"body,omitempty"`
}

func (o UpdateInstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceGroupRequest", string(data)}, " ")
}
