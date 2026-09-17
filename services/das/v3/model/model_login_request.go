package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginRequest Request Object
type LoginRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o LoginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginRequest struct{}"
	}

	return strings.Join([]string{"LoginRequest", string(data)}, " ")
}
