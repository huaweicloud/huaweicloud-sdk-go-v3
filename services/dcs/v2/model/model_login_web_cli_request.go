package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginWebCliRequest Request Object
type LoginWebCliRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *LoginWebCliBody `json:"body,omitempty"`
}

func (o LoginWebCliRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWebCliRequest struct{}"
	}

	return strings.Join([]string{"LoginWebCliRequest", string(data)}, " ")
}
