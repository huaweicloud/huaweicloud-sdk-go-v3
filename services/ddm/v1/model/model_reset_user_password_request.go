package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResetUserPasswordRequest struct {
	// DDM实例ID。

	InstanceId string `json:"instance_id"`
	// 需要修改的DDM帐号名称。

	Username string `json:"username"`

	Body *ResetUserPasswordReq `json:"body,omitempty"`
}

func (o ResetUserPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordRequest", string(data)}, " ")
}
