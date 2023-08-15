package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPasswordRequest Request Object
type ResetPasswordRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResetInstancePasswordBody `json:"body,omitempty"`
}

func (o ResetPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequest", string(data)}, " ")
}
