package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroupRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// 安全组ID列表(目前只支持传一个ID)
	SecuritygroupIds []string `json:"securitygroup_ids"`
}

func (o SecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"SecurityGroupRequest", string(data)}, " ")
}
