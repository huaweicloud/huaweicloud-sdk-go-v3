package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicIpRequest Request Object
type DeletePublicIpRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeletePublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicIpRequest struct{}"
	}

	return strings.Join([]string{"DeletePublicIpRequest", string(data)}, " ")
}
