package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LoadbalancerResp 负载均衡器响应体
type LoadbalancerResp struct {

	// 负载均衡器ID
	Id string `json:"id"`

	// 负载均衡器所在的项目ID。
	TenantId string `json:"tenant_id"`

	// 负载均衡器名称。
	Name string `json:"name"`

	// 负载均衡器的描述信息
	Description string `json:"description"`

	// 负载均衡器所在的子网的IPv4子网ID。仅支持内网类型。
	VipSubnetId string `json:"vip_subnet_id"`

	// 负载均衡器虚拟IP对应的端口ID
	VipPortId string `json:"vip_port_id"`

	// 负载均衡器的虚拟IP。
	VipAddress string `json:"vip_address"`

	// 负载均衡器关联的监听器ID的列表
	Listeners []ResourceList `json:"listeners"`

	// 负载均衡器关联的后端云服务器组ID的列表。
	Pools []ResourceList `json:"pools"`

	// 负载均衡器的供应者名称。只支持vlb
	Provider string `json:"provider"`

	// 负载均衡器的操作状态
	OperatingStatus LoadbalancerRespOperatingStatus `json:"operating_status"`

	// 负载均衡器的配置状态
	ProvisioningStatus LoadbalancerRespProvisioningStatus `json:"provisioning_status"`

	// 负载均衡器的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp bool `json:"admin_state_up"`

	// 负载均衡器的创建时间
	CreatedAt string `json:"created_at"`

	// 负载均衡器的更新时间
	UpdatedAt string `json:"updated_at"`

	// 负载均衡器的企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 负载均衡器所在的项目ID。
	ProjectId string `json:"project_id"`

	// 负载均衡器的标签列表
	Tags []string `json:"tags"`

	// 负载均衡器绑定的公网IP。只支持绑定一个公网IP。
	Publicips []PublicIpInfo `json:"publicips"`

	// 收费模式。取值：  flavor：按规格计费 lcu：按使用量计费 说明：弹性扩缩容实例该字段无效，按lcu收费；包周期实例该字段无效，预付费收费。
	ChargeMode string `json:"charge_mode"`

	// 资源账单信息，取值：     - 空：按需计费。     - 非空：包周期计费，  包周期计费billing_info字段的格式为：order_id:product_id:region_id:project_id。
	BillingInfo string `json:"billing_info"`

	// 负载均衡器的冻结场景。若负载均衡器有多个冻结场景，用逗号分隔。取值：  POLICE：公安冻结场景。 ILLEGAL：违规冻结场景。 VERIFY：客户未实名认证冻结场景。 PARTNER：合作伙伴冻结（合作伙伴冻结子客户资源）。 AREAR：欠费冻结场景。
	FrozenScene *string `json:"frozen_scene,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *LoadbalancerRespProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o LoadbalancerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadbalancerResp struct{}"
	}

	return strings.Join([]string{"LoadbalancerResp", string(data)}, " ")
}

type LoadbalancerRespOperatingStatus struct {
	value string
}

type LoadbalancerRespOperatingStatusEnum struct {
	ONLINE     LoadbalancerRespOperatingStatus
	OFFLINE    LoadbalancerRespOperatingStatus
	DEGRADED   LoadbalancerRespOperatingStatus
	DISABLED   LoadbalancerRespOperatingStatus
	NO_MONITOR LoadbalancerRespOperatingStatus
}

func GetLoadbalancerRespOperatingStatusEnum() LoadbalancerRespOperatingStatusEnum {
	return LoadbalancerRespOperatingStatusEnum{
		ONLINE: LoadbalancerRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: LoadbalancerRespOperatingStatus{
			value: "OFFLINE",
		},
		DEGRADED: LoadbalancerRespOperatingStatus{
			value: "DEGRADED",
		},
		DISABLED: LoadbalancerRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: LoadbalancerRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c LoadbalancerRespOperatingStatus) Value() string {
	return c.value
}

func (c LoadbalancerRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespOperatingStatus) UnmarshalJSON(b []byte) error {
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

type LoadbalancerRespProvisioningStatus struct {
	value string
}

type LoadbalancerRespProvisioningStatusEnum struct {
	ACTIVE         LoadbalancerRespProvisioningStatus
	PENDING_CREATE LoadbalancerRespProvisioningStatus
	ERROR          LoadbalancerRespProvisioningStatus
}

func GetLoadbalancerRespProvisioningStatusEnum() LoadbalancerRespProvisioningStatusEnum {
	return LoadbalancerRespProvisioningStatusEnum{
		ACTIVE: LoadbalancerRespProvisioningStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: LoadbalancerRespProvisioningStatus{
			value: "PENDING_CREATE",
		},
		ERROR: LoadbalancerRespProvisioningStatus{
			value: "ERROR",
		},
	}
}

func (c LoadbalancerRespProvisioningStatus) Value() string {
	return c.value
}

func (c LoadbalancerRespProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespProvisioningStatus) UnmarshalJSON(b []byte) error {
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

type LoadbalancerRespProtectionStatus struct {
	value string
}

type LoadbalancerRespProtectionStatusEnum struct {
	NON_PROTECTION     LoadbalancerRespProtectionStatus
	CONSOLE_PROTECTION LoadbalancerRespProtectionStatus
}

func GetLoadbalancerRespProtectionStatusEnum() LoadbalancerRespProtectionStatusEnum {
	return LoadbalancerRespProtectionStatusEnum{
		NON_PROTECTION: LoadbalancerRespProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: LoadbalancerRespProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c LoadbalancerRespProtectionStatus) Value() string {
	return c.value
}

func (c LoadbalancerRespProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespProtectionStatus) UnmarshalJSON(b []byte) error {
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
