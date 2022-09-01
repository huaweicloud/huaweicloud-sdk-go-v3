package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceSecurityGroupRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ModifyInstanceSecurityGroupReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateInstanceSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceSecurityGroupRequest", string(data)}, " ")
}
