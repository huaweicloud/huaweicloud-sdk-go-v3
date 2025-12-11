package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationNodeResponseInfo 组织结构树
type OrganizationNodeResponseInfo struct {

	// **参数解释** 当前组织节点的父节点唯一标识ID，用于标识组织树层级关系（根节点父ID通常为“0”或空）； **取值范围** 字符长度1-64位，符合平台组织节点ID命名规范
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释** 组织树节点的唯一标识ID，用于唯一确定某个组织节点； **取值范围** 字符长度1-64位，符合平台组织节点ID命名规范（如UUID或数字组合）
	Id *string `json:"id,omitempty"`

	// **参数解释** 组织节点的统一资源名称（URN），用于跨服务唯一标识组织资源； **取值范围** 字符长度1-256位，格式为organizations::{management_account_id}:xxxxx:{org_id}/xxxxxxxx，符合平台URN命名规范
	Urn *string `json:"urn,omitempty"`

	// **参数解释** 组织节点的名称（可能是组织单元名称或账号名称，与org_type对应）； **取值范围** 字符长度1-64位，支持字母、数字、连字符、下划线及中文，不能以特殊字符开头或结尾
	Name *string `json:"name,omitempty"`

	// **参数解释**: 节点（组织单元或账号）的类型； **取值范围**: unit（组织单元）、account（账号）
	OrgType *string `json:"org_type,omitempty"`

	// **参数解释**: 标识组织节点（组织单元或账号）是否已完成授权； **取值范围**: true（已授权，无需重复授权）、false（未授权，需完成授权后使用）
	Delegated *bool `json:"delegated,omitempty"`
}

func (o OrganizationNodeResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationNodeResponseInfo struct{}"
	}

	return strings.Join([]string{"OrganizationNodeResponseInfo", string(data)}, " ")
}
