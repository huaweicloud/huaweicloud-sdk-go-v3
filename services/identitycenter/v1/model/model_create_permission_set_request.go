package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermissionSetRequest Request Object
type CreatePermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	Body *CreatePermissionSetReqBody `json:"body,omitempty"`
}

func (o CreatePermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermissionSetRequest struct{}"
	}

	return strings.Join([]string{"CreatePermissionSetRequest", string(data)}, " ")
}
