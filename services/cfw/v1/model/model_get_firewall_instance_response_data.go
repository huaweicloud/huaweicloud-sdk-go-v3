package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GetFirewallInstanceResponseData struct {

	// 防火墙id
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 防火墙名称
	Name *string `json:"name,omitempty"`

	// 集群类型
	HaType *int32 `json:"ha_type,omitempty"`

	// 计费模式 0：包年/包月 1：按需
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	// 防护对象列表
	ProtectObjects *[]ProtectObjectVo `json:"protect_objects,omitempty"`

	// 防火墙状态列表，包括-1：等待支付，0：创建中，1，删除中，2：运行中，3：升级中，4：删除完成：5：冻结中，6：创建失败，7：删除失败，8：冻结失败，9：存储中，10：存储失败，11：升级失败
	Status *GetFirewallInstanceResponseDataStatus `json:"status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否为旧引擎，true表示是，false表示不是
	IsOldFirewallInstance *bool `json:"is_old_firewall_instance,omitempty"`

	// 是否支持ipv6，true表示是，false表示不是
	SupportIpv6 *bool `json:"support_ipv6,omitempty"`

	// 特性开关，boolean值为true表示是，false表示否
	FeatureToggle map[string]bool `json:"feature_toggle,omitempty"`
}

func (o GetFirewallInstanceResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFirewallInstanceResponseData struct{}"
	}

	return strings.Join([]string{"GetFirewallInstanceResponseData", string(data)}, " ")
}

type GetFirewallInstanceResponseDataStatus struct {
	value int32
}

type GetFirewallInstanceResponseDataStatusEnum struct {
	E_MINUS_1 GetFirewallInstanceResponseDataStatus
	E_0       GetFirewallInstanceResponseDataStatus
	E_1       GetFirewallInstanceResponseDataStatus
	E_2       GetFirewallInstanceResponseDataStatus
	E_3       GetFirewallInstanceResponseDataStatus
	E_4       GetFirewallInstanceResponseDataStatus
	E_5       GetFirewallInstanceResponseDataStatus
	E_6       GetFirewallInstanceResponseDataStatus
	E_7       GetFirewallInstanceResponseDataStatus
	E_8       GetFirewallInstanceResponseDataStatus
	E_9       GetFirewallInstanceResponseDataStatus
	E_10      GetFirewallInstanceResponseDataStatus
	E_11      GetFirewallInstanceResponseDataStatus
}

func GetGetFirewallInstanceResponseDataStatusEnum() GetFirewallInstanceResponseDataStatusEnum {
	return GetFirewallInstanceResponseDataStatusEnum{
		E_MINUS_1: GetFirewallInstanceResponseDataStatus{
			value: -1,
		}, E_0: GetFirewallInstanceResponseDataStatus{
			value: 0,
		}, E_1: GetFirewallInstanceResponseDataStatus{
			value: 1,
		}, E_2: GetFirewallInstanceResponseDataStatus{
			value: 2,
		}, E_3: GetFirewallInstanceResponseDataStatus{
			value: 3,
		}, E_4: GetFirewallInstanceResponseDataStatus{
			value: 4,
		}, E_5: GetFirewallInstanceResponseDataStatus{
			value: 5,
		}, E_6: GetFirewallInstanceResponseDataStatus{
			value: 6,
		}, E_7: GetFirewallInstanceResponseDataStatus{
			value: 7,
		}, E_8: GetFirewallInstanceResponseDataStatus{
			value: 8,
		}, E_9: GetFirewallInstanceResponseDataStatus{
			value: 9,
		}, E_10: GetFirewallInstanceResponseDataStatus{
			value: 10,
		}, E_11: GetFirewallInstanceResponseDataStatus{
			value: 11,
		},
	}
}

func (c GetFirewallInstanceResponseDataStatus) Value() int32 {
	return c.value
}

func (c GetFirewallInstanceResponseDataStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetFirewallInstanceResponseDataStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
