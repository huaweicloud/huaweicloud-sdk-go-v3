package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 存储库订单
type VaultOrder struct {

	// 存储库名称  最小长度：1  最大长度：64
	Name *string `json:"name,omitempty"`

	Billing *BillingCreate `json:"billing"`

	// 绑定的备份资源，未在创建时绑定资源填[]
	Resources []ResourceCreate `json:"resources"`

	// 描述  最小长度：0  最大长度：255
	Description *string `json:"description,omitempty"`

	// 备份策略ID，不设置时为null，不自动备份。
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	// 标签列表 tags不允许为空列表。 tags中最多包含10个key。 tags中key不允许重复。
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目ID，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否支持自动挂载。
	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	// 存储库阈值，百分比。  最小值：1  最大值：100
	Threshold *int32 `json:"threshold,omitempty"`

	// 当容量到达阈值，是否启用通知
	SmnNotify *bool `json:"smn_notify,omitempty"`

	Parameters *VaultCreateParameters `json:"parameters,omitempty"`

	// 是否开启存储库自动扩容能力（只支持按需存储库）。
	AutoExpand *bool `json:"auto_expand,omitempty"`
}

func (o VaultOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultOrder struct{}"
	}

	return strings.Join([]string{"VaultOrder", string(data)}, " ")
}
