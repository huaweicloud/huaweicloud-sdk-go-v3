package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可恢复的实例信息结构体
type QueryRestoreList struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 实例模式
	InstanceMode string `json:"instance_mode"`

	// 引擎名称
	EngineName string `json:"engine_name"`

	// 引擎版本
	EngineVersion string `json:"engine_version"`

	// VPC ID。
	VpcId string `json:"vpc_id"`

	// 子网ID列表
	SubnetIds []string `json:"subnet_ids"`

	// 安全组ID列表
	SecurityGroupIds []string `json:"security_group_ids"`
}

func (o QueryRestoreList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryRestoreList struct{}"
	}

	return strings.Join([]string{"QueryRestoreList", string(data)}, " ")
}
