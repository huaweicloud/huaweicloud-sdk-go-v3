package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAdministratorRequest Request Object
type ResetAdministratorRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	Body *AdminUserInfoReq `json:"body,omitempty"`
}

func (o ResetAdministratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAdministratorRequest struct{}"
	}

	return strings.Join([]string{"ResetAdministratorRequest", string(data)}, " ")
}
