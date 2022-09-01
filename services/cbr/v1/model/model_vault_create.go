package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultCreate struct {

	// 备份策略ID，不设置时为null，不自动备份。
	BackupPolicyId *string `json:"backup_policy_id,omitempty" xml:"backup_policy_id"`

	Billing *BillingCreate `json:"billing" xml:"billing"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 存储库名称
	Name string `json:"name" xml:"name"`

	// 绑定的备份资源，未在创建时绑定资源填[]
	Resources []ResourceCreate `json:"resources" xml:"resources"`

	// 标签列表 tags不允许为空列表。 tags中最多包含10个key。 tags中key不允许重复。
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 企业项目ID，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 是否支持自动挂载。
	AutoBind *bool `json:"auto_bind,omitempty" xml:"auto_bind"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty" xml:"bind_rules"`

	// 是否自动扩容。按需存储库支持自动扩容，包周期存储库不支持扩容。
	AutoExpand *bool `json:"auto_expand,omitempty" xml:"auto_expand"`
}

func (o VaultCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreate struct{}"
	}

	return strings.Join([]string{"VaultCreate", string(data)}, " ")
}
