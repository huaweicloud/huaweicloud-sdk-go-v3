package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnbindApiForAcl struct {
	// API的认证方式

	AuthType *string `json:"auth_type,omitempty"`
	// 发布的环境名

	RunEnvName *string `json:"run_env_name,omitempty"`
	// API所属分组的名称

	GroupName *string `json:"group_name,omitempty"`
	// 发布记录的编号

	PublishId *string `json:"publish_id,omitempty"`
	// API所属分组的编号

	GroupId *string `json:"group_id,omitempty"`
	// API名称

	Name *string `json:"name,omitempty"`
	// API描述

	Remark *string `json:"remark,omitempty"`
	// 发布的环境id

	RunEnvId *string `json:"run_env_id,omitempty"`
	// API编号

	Id *string `json:"id,omitempty"`
	// API的请求地址

	ReqUri *string `json:"req_uri,omitempty"`
	// API类型

	Type *int32 `json:"type,omitempty"`
	// 绑定的其他同类型的ACL策略名称

	AclName *string `json:"acl_name,omitempty"`
}

func (o UnbindApiForAcl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindApiForAcl struct{}"
	}

	return strings.Join([]string{"UnbindApiForAcl", string(data)}, " ")
}
