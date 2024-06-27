package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetFirewallInstanceResponseRecord struct {

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防火墙名称
	Name *string `json:"name,omitempty"`

	// 集群类型
	HaType *int32 `json:"ha_type,omitempty"`

	// 计费模式 0：包年/包月 1：按需
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 引擎类型
	EngineType *int32 `json:"engine_type,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// 防护对象列表
	ProtectObjects *[]ProtectObjectVo `json:"protect_objects,omitempty"`

	// 防火墙状态列表，包括-1：等待支付，0：创建中，1，删除中，2：运行中，3：升级中，4：删除完成：5：冻结中，6：创建失败，7：删除失败，8：冻结失败，9：存储中，10：存储失败，11：升级失败
	Status *GetFirewallInstanceResponseRecordStatus `json:"status,omitempty"`

	// 是否为旧引擎，true表示是，false表示不是
	IsOldFirewallInstance *bool `json:"is_old_firewall_instance,omitempty"`

	// 是否支持obs
	IsAvailableObs *bool `json:"is_available_obs,omitempty"`

	// 是否支持威胁标签
	IsSupportThreatTags *bool `json:"is_support_threat_tags,omitempty"`

	// 是否支持ipv6，true表示是，false表示不是
	SupportIpv6 *bool `json:"support_ipv6,omitempty"`

	// 特性开关，boolean值为true表示是，false表示否
	FeatureToggle map[string]bool `json:"feature_toggle,omitempty"`

	// 防火墙资源列表
	Resources *[]FirewallInstanceResource `json:"resources,omitempty"`

	// 防火墙名称
	FwInstanceName *string `json:"fw_instance_name,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 是否支持url过滤，true表示是，false表示不是
	SupportUrlFiltering *bool `json:"support_url_filtering,omitempty"`

	// 标签列表
	Tags *string `json:"tags,omitempty"`
}

func (o GetFirewallInstanceResponseRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFirewallInstanceResponseRecord struct{}"
	}

	return strings.Join([]string{"GetFirewallInstanceResponseRecord", string(data)}, " ")
}

type GetFirewallInstanceResponseRecordStatus struct {
	value int32
}

type GetFirewallInstanceResponseRecordStatusEnum struct {
	E_MINUS_1 GetFirewallInstanceResponseRecordStatus
	E_0       GetFirewallInstanceResponseRecordStatus
	E_1       GetFirewallInstanceResponseRecordStatus
	E_2       GetFirewallInstanceResponseRecordStatus
	E_3       GetFirewallInstanceResponseRecordStatus
	E_4       GetFirewallInstanceResponseRecordStatus
	E_5       GetFirewallInstanceResponseRecordStatus
	E_6       GetFirewallInstanceResponseRecordStatus
	E_7       GetFirewallInstanceResponseRecordStatus
	E_8       GetFirewallInstanceResponseRecordStatus
	E_9       GetFirewallInstanceResponseRecordStatus
	E_10      GetFirewallInstanceResponseRecordStatus
	E_11      GetFirewallInstanceResponseRecordStatus
}

func GetGetFirewallInstanceResponseRecordStatusEnum() GetFirewallInstanceResponseRecordStatusEnum {
	return GetFirewallInstanceResponseRecordStatusEnum{
		E_MINUS_1: GetFirewallInstanceResponseRecordStatus{
			value: -1,
		}, E_0: GetFirewallInstanceResponseRecordStatus{
			value: 0,
		}, E_1: GetFirewallInstanceResponseRecordStatus{
			value: 1,
		}, E_2: GetFirewallInstanceResponseRecordStatus{
			value: 2,
		}, E_3: GetFirewallInstanceResponseRecordStatus{
			value: 3,
		}, E_4: GetFirewallInstanceResponseRecordStatus{
			value: 4,
		}, E_5: GetFirewallInstanceResponseRecordStatus{
			value: 5,
		}, E_6: GetFirewallInstanceResponseRecordStatus{
			value: 6,
		}, E_7: GetFirewallInstanceResponseRecordStatus{
			value: 7,
		}, E_8: GetFirewallInstanceResponseRecordStatus{
			value: 8,
		}, E_9: GetFirewallInstanceResponseRecordStatus{
			value: 9,
		}, E_10: GetFirewallInstanceResponseRecordStatus{
			value: 10,
		}, E_11: GetFirewallInstanceResponseRecordStatus{
			value: 11,
		},
	}
}

func (c GetFirewallInstanceResponseRecordStatus) Value() int32 {
	return c.value
}

func (c GetFirewallInstanceResponseRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetFirewallInstanceResponseRecordStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
