package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbrVaultDetailResourceDetail 资源详情
type CbrVaultDetailResourceDetail struct {

	// 规格编码。
	SpecCode *string `json:"spec_code,omitempty"`

	// 是否跨账号
	CrossAccount *bool `json:"cross_account,omitempty"`

	// 创建模式，按需：post_paid，包周期：pre_paid，默认为post_paid
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 对象类型：云服务器（server），云硬盘（disk），文件系统（turbo），云桌面（workspace），VMware（vmware），关系型数据库（rds），文件（file）
	ObjectType *string `json:"object_type,omitempty"`

	// 保护类型：备份（backup）、复制(replication)。
	ProtectType *string `json:"protect_type,omitempty"`

	// 存储库关联的保护策略id
	PolicyIds *[]string `json:"policy_ids,omitempty"`

	// 已使用容量，单位MB
	Used *string `json:"used,omitempty"`

	// 绑定规则
	BindRules *string `json:"bind_rules,omitempty"`

	// 容量，单位GB
	Size *int32 `json:"size,omitempty"`

	// 存储库资源列表
	VaultResources *[]string `json:"vault_resources,omitempty"`

	// 跨账号URN
	CrossAccountUrn *string `json:"cross_account_urn,omitempty"`

	// 存储库资源类型id
	ProviderId *string `json:"provider_id,omitempty"`

	// 是否已锁定
	Locked *bool `json:"locked,omitempty"`

	// 是否开启存储库自动扩容能力（只支持按需存储库）。
	AutoExpand *bool `json:"auto_expand,omitempty"`

	// 存储库是否多az
	IsMultiAz *bool `json:"is_multi_az,omitempty"`

	// 保护状态
	ProtectStatus *[]string `json:"protect_status,omitempty"`

	// status
	Status *string `json:"status,omitempty"`
}

func (o CbrVaultDetailResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbrVaultDetailResourceDetail struct{}"
	}

	return strings.Join([]string{"CbrVaultDetailResourceDetail", string(data)}, " ")
}
