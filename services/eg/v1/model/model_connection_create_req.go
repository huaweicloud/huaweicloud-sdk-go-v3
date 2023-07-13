package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionCreateReq struct {

	// 目标连接名称，租户下唯一，由字母、数字、点、下划线和中划线组成，必须以字母或数字开头，不能为default
	Name string `json:"name"`

	// 目标连接描述
	Description *string `json:"description,omitempty"`

	// 待连接的VPC ID
	VpcId string `json:"vpc_id"`

	// 待连接的子网ID
	SubnetId string `json:"subnet_id"`

	Type *ConnectionType `json:"type,omitempty"`

	KafkaDetail *KafkaConnectionDetail `json:"kafka_detail,omitempty"`
}

func (o ConnectionCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionCreateReq struct{}"
	}

	return strings.Join([]string{"ConnectionCreateReq", string(data)}, " ")
}
