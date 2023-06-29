package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRequest Request Object
type UpdateUserRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	// 需要修改的DDM帐号名称。
	Username string `json:"username"`

	Body *UpdateUserReq `json:"body,omitempty"`
}

func (o UpdateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserRequest", string(data)}, " ")
}
