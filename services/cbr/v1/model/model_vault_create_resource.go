package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultCreateResource struct {
	Billing *Billing `json:"billing"`

	// 存储库自定义描述信息。
	Description *string `json:"description,omitempty"`

	// 存储库ID
	Id string `json:"id"`

	// 存储库名称
	Name string `json:"name"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 存储库资源类型id
	ProviderId string `json:"provider_id"`

	// 存储库资源
	Resources []ResourceResp `json:"resources"`

	// 存储库标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目id，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否自动绑定，默认为false，不支持。
	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 创建时间,例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *string `json:"created_at,omitempty"`

	// [是否开启存储库自动扩容能力（只支持按需存储库）。](tag:hws,hws_hk) [是否开启存储库自动扩容能力。](tag:dt,ocb,tlf,sbc,fcs_vm,ctc,g42,tm,cmcc,tm,hcso_dt)
	AutoExpand *bool `json:"auto_expand,omitempty"`

	// 存储库smn消息通知开关
	SmnNotify *bool `json:"smn_notify,omitempty"`

	// 存储库容量阈值，已用容量占总容量达到此百分比即发送相关通知
	Threshold *int32 `json:"threshold,omitempty"`

	// 包周期创建错误信息
	ErrText *string `json:"errText,omitempty"`

	// 包周期订购结果
	RetCode *string `json:"retCode,omitempty"`

	// 包周期创建订单信息
	Orders *[]CbcOrderResult `json:"orders,omitempty"`

	// 备份名称前缀
	BackupNamePrefix *string `json:"backup_name_prefix,omitempty"`

	// 是否允许使用超出存储库容量
	DemandBilling *bool `json:"demand_billing,omitempty"`

	// 存储库删除次数
	CbcDeleteCount *int32 `json:"cbc_delete_count,omitempty"`

	// 存储库是否冻结
	Frozen *bool `json:"frozen,omitempty"`

	// 用于标识SMB服务
	SysLockSourceService *string `json:"sys_lock_source_service,omitempty"`

	// 用于标识该存储库是否已锁定
	Locked *bool `json:"locked,omitempty"`
}

func (o VaultCreateResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreateResource struct{}"
	}

	return strings.Join([]string{"VaultCreateResource", string(data)}, " ")
}
