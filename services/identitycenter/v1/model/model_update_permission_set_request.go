package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePermissionSetRequest Request Object
type UpdatePermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`

	Body *UpdatePermissionSetReqBody `json:"body,omitempty"`
}

func (o UpdatePermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionSetRequest struct{}"
	}

	return strings.Join([]string{"UpdatePermissionSetRequest", string(data)}, " ")
}
