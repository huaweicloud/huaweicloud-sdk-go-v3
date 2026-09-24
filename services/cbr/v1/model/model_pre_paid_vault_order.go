package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrePaidVaultOrder 包周期存储库订单
type PrePaidVaultOrder struct {

	// 存储库名称，最大支持64字符，只能由中文、字母、数字、\"_\"、\"-\"组成。默认取值不涉及。
	Name *string `json:"name,omitempty"`

	Billing *PrePaidBillingCreate `json:"billing"`

	// 绑定的备份资源，未在创建时绑定资源填[]
	Resources []ResourceCreate `json:"resources"`

	// 存储库描述，取值范围：最小长度：0，最大长度：255。默认取值不涉及。
	Description *string `json:"description,omitempty"`

	// 备份策略ID，默认值为null，不自动备份。 [获取方法请参见\"[获取备份策略ID](https://support.huaweicloud.com/api-cbr/ListPolicies.html)\"。](tag:hws) [获取方法请参见\"[获取备份策略ID](https://support.huaweicloud.com/intl/zh-cn/api-cbr/ListPolicies.html)\"。](tag:hws_hk)
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	// 标签列表 tags不允许为空列表。 tags中最多包含10个key。 tags中key不允许重复。
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目ID，默认为'0'。 [获取方法请参见\"[获取企业项目ID](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)\"。](tag:hws) [获取方法请参见\"[获取企业项目ID](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/zh-cn_topic_0126101490.html)\"。](tag:hws_hk)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 功能说明：是否支持自动挂载。默认为false。 取值范围： - true：支持自动挂载 - false：不支持自动挂载
	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	// 功能说明：存储库容量阈值，存储库已用容量和总容量的百分比超过该值，如果smn_notify为开，将发送相关通知。 取值范围：[1, 100]，默认值为80。
	Threshold *int32 `json:"threshold,omitempty"`

	// 功能说明：是否发送smn通知开关，默认为true 取值范围： - true：发送smn通知 - false：不发送smn通知
	SmnNotify *bool `json:"smn_notify,omitempty"`

	Parameters *VaultCreateParameters `json:"parameters,omitempty"`

	// 功能说明：是否开启存储库自动扩容能力（只支持按需存储库），默认为false。 取值范围： - true：支持自动扩容； - false：不支持自动扩容。
	AutoExpand *bool `json:"auto_expand,omitempty"`

	// 功能说明：用于标识当前存储库是否已锁定，锁定的存储库不支持解锁。默认值为false。 [关于备份锁定的详细信息，请参考\"[开启备份锁定](https://support.huaweicloud.com/usermanual-cbr/cbr_01_0035.html)\"。](tag:hws) [关于备份锁定的详细信息，请参考\"[开启备份锁定](https://support.huaweicloud.com/intl/zh-cn/usermanual-cbr/cbr_01_0035.html)\"。](tag:hws_hk) 取值范围： - true：锁定存储库 - false：不锁定存储库
	Locked *bool `json:"locked,omitempty"`

	// 功能说明：是否为跨账号复制存储库，默认值为false，只有创建跨账号复制存储库时才允许该值为true。 取值范围： - false: 非跨账号复制存储库 - true: 跨账号复制存储库
	CrossAccount *bool `json:"cross_account,omitempty"`

	DataEncryption *DataEncryption `json:"data_encryption,omitempty"`
}

func (o PrePaidVaultOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidVaultOrder struct{}"
	}

	return strings.Join([]string{"PrePaidVaultOrder", string(data)}, " ")
}
