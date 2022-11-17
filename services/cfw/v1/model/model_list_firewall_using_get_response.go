package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListFirewallUsingGetResponse struct {

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
	Status *ListFirewallUsingGetResponseStatus `json:"status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否为旧引擎，true表示是，false表示不是
	IsOldFirewallInstance *bool `json:"is_old_firewall_instance,omitempty"`

	// 是否支持ipv6，true表示是，false表示不是
	SupportIpv6 *bool `json:"support_ipv6,omitempty"`

	// 特性开关，boolean值为true表示是，false表示否
	FeatureToggle  map[string]bool `json:"feature_toggle,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListFirewallUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListFirewallUsingGetResponse", string(data)}, " ")
}

type ListFirewallUsingGetResponseStatus struct {
	value int32
}

type ListFirewallUsingGetResponseStatusEnum struct {
	E_MINUS_1 ListFirewallUsingGetResponseStatus
	E_0       ListFirewallUsingGetResponseStatus
	E_1       ListFirewallUsingGetResponseStatus
	E_2       ListFirewallUsingGetResponseStatus
	E_3       ListFirewallUsingGetResponseStatus
	E_4       ListFirewallUsingGetResponseStatus
	E_5       ListFirewallUsingGetResponseStatus
	E_6       ListFirewallUsingGetResponseStatus
	E_7       ListFirewallUsingGetResponseStatus
	E_8       ListFirewallUsingGetResponseStatus
	E_9       ListFirewallUsingGetResponseStatus
	E_10      ListFirewallUsingGetResponseStatus
	E_11      ListFirewallUsingGetResponseStatus
}

func GetListFirewallUsingGetResponseStatusEnum() ListFirewallUsingGetResponseStatusEnum {
	return ListFirewallUsingGetResponseStatusEnum{
		E_MINUS_1: ListFirewallUsingGetResponseStatus{
			value: -1,
		}, E_0: ListFirewallUsingGetResponseStatus{
			value: 0,
		}, E_1: ListFirewallUsingGetResponseStatus{
			value: 1,
		}, E_2: ListFirewallUsingGetResponseStatus{
			value: 2,
		}, E_3: ListFirewallUsingGetResponseStatus{
			value: 3,
		}, E_4: ListFirewallUsingGetResponseStatus{
			value: 4,
		}, E_5: ListFirewallUsingGetResponseStatus{
			value: 5,
		}, E_6: ListFirewallUsingGetResponseStatus{
			value: 6,
		}, E_7: ListFirewallUsingGetResponseStatus{
			value: 7,
		}, E_8: ListFirewallUsingGetResponseStatus{
			value: 8,
		}, E_9: ListFirewallUsingGetResponseStatus{
			value: 9,
		}, E_10: ListFirewallUsingGetResponseStatus{
			value: 10,
		}, E_11: ListFirewallUsingGetResponseStatus{
			value: 11,
		},
	}
}

func (c ListFirewallUsingGetResponseStatus) Value() int32 {
	return c.value
}

func (c ListFirewallUsingGetResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFirewallUsingGetResponseStatus) UnmarshalJSON(b []byte) error {
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
