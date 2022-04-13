package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultCreate struct {
	// 备份策略ID，不设置时为null，不自动备份。

	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	Billing *BillingCreate `json:"billing"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 存储库名称

	Name string `json:"name"`
	// 绑定的备份资源，未在创建时绑定资源填[]

	Resources []ResourceResp `json:"resources"`
	// 标签列表 tags不允许为空列表。 tags中最多包含10个key。 tags中key不允许重复。

	Tags *[]Tag `json:"tags,omitempty"`
	// 企业项目ID，默认为‘0’。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 是否支持自动挂载。

	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`
	// 是否自动扩容。按需存储库支持自动扩容，包周期存储库不支持扩容。

	AutoExpand *bool `json:"auto_expand,omitempty"`
}

func (o VaultCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreate struct{}"
	}

	return strings.Join([]string{"VaultCreate", string(data)}, " ")
}
