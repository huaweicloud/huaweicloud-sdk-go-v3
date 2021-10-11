package model

import (
	"encoding/json"

	"strings"
)

type UnbindApiForAclPageApisResp struct {
	// API的ID

	Id *string `json:"id,omitempty"`
	// API名称

	Name *string `json:"name,omitempty"`
	// API所属分组的编号

	GroupId *string `json:"group_id,omitempty"`
	// API所属分组的名称

	GroupName *string `json:"group_name,omitempty"`
	// API开放状态

	Type *int32 `json:"type,omitempty"`
	// API描述

	Remark *string `json:"remark,omitempty"`
	// 发布的环境名

	RunEnvName *string `json:"run_env_name,omitempty"`
	// 发布的环境id

	RunEnvId *string `json:"run_env_id,omitempty"`
	// API发布记录编号

	PublishId *string `json:"publish_id,omitempty"`
	// 绑定的其他同类型的ACL策略名称

	AclName *string `json:"acl_name,omitempty"`
}

func (o UnbindApiForAclPageApisResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnbindApiForAclPageApisResp struct{}"
	}

	return strings.Join([]string{"UnbindApiForAclPageApisResp", string(data)}, " ")
}
