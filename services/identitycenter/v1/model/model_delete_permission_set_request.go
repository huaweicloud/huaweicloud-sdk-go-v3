package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePermissionSetRequest Request Object
type DeletePermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`
}

func (o DeletePermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePermissionSetRequest struct{}"
	}

	return strings.Join([]string{"DeletePermissionSetRequest", string(data)}, " ")
}
