package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceId 网络实例的ID。
type InstanceId struct {

	// 网络实例（VPC，VGW）的ID。
	InstanceId string `json:"instance_id"`
}

func (o InstanceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceId struct{}"
	}

	return strings.Join([]string{"InstanceId", string(data)}, " ")
}
