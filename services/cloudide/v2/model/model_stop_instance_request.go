package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopInstanceRequest Request Object
type StopInstanceRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o StopInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceRequest", string(data)}, " ")
}
