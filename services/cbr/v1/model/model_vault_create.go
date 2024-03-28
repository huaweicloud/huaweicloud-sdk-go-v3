package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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
	Resources []ResourceCreate `json:"resources"`

	// 标签列表 tags不允许为空列表。 tags中最多包含10个key。 tags中key不允许重复。
	Tags *[]Tag `json:"tags,omitempty"`

	// 企业项目ID，默认为‘0’。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否支持自动挂载。
	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	// [是否开启存储库自动扩容能力（只支持按需存储库）。](tag:hws,hws_hk) [是否开启存储库自动扩容能力。](tag:dt,ocb,tlf,sbc,fcs_vm,ctc,g42,tm,cmcc,hcso_dt)
	AutoExpand *bool `json:"auto_expand,omitempty"`

	// 存储库容量阈值，已用容量占总容量达到此百分比，将根据 smn_notify 参数设置选择是否发送相关通知。 默认值为：80 最大值：100 最小值：1
	Threshold *int32 `json:"threshold,omitempty"`

	// 存储库smn消息通知开关。 默认值为 true。
	SmnNotify *bool `json:"smn_notify,omitempty"`

	// 备份名称前缀，设置后该存储库自动备份产生的备份副本都将携带该备份名称前缀
	BackupNamePrefix *string `json:"backup_name_prefix,omitempty"`

	// 存储库使用是否允许超出容量，只有创建包周期存储库时才允许该值为 true
	DemandBilling *bool `json:"demand_billing,omitempty"`

	// 用于标识SMB服务，您可以设置为SMB或者空
	SysLockSourceService *VaultCreateSysLockSourceService `json:"sys_lock_source_service,omitempty"`
}

func (o VaultCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultCreate struct{}"
	}

	return strings.Join([]string{"VaultCreate", string(data)}, " ")
}

type VaultCreateSysLockSourceService struct {
	value string
}

type VaultCreateSysLockSourceServiceEnum struct {
	SMB   VaultCreateSysLockSourceService
	EMPTY VaultCreateSysLockSourceService
}

func GetVaultCreateSysLockSourceServiceEnum() VaultCreateSysLockSourceServiceEnum {
	return VaultCreateSysLockSourceServiceEnum{
		SMB: VaultCreateSysLockSourceService{
			value: "SMB",
		},
		EMPTY: VaultCreateSysLockSourceService{
			value: "",
		},
	}
}

func (c VaultCreateSysLockSourceService) Value() string {
	return c.value
}

func (c VaultCreateSysLockSourceService) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VaultCreateSysLockSourceService) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
