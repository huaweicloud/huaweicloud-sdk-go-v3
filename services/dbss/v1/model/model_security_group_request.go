package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroupRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

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
