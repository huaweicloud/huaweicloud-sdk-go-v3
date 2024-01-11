package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateLoadbalancerReq 创建负载均衡器的请求体
type CreateLoadbalancerReq struct {

	// 负载均衡器所在的项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 负载均衡器名称。
	Name *string `json:"name,omitempty"`

	// 负载均衡器的描述信息
	Description *string `json:"description,omitempty"`

	// 负载均衡器所在的子网IPv4子网ID
	VipSubnetId string `json:"vip_subnet_id"`

	// 负载均衡器的虚拟IP。
	VipAddress *string `json:"vip_address,omitempty"`

	// 负载均衡器的供应者名称。只支持vlb
	Provider *CreateLoadbalancerReqProvider `json:"provider,omitempty"`

	// 负载均衡器的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 企业项目ID。创建负载均衡器时，给负载均衡器绑定企业项目ID。取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 关于企业项目ID的获取及企业项目特性的详细信息，请参见《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *CreateLoadbalancerReqProtectionStatus `json:"protection_status,omitempty"`

	// 设置保护的原因 >仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o CreateLoadbalancerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerReq struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerReq", string(data)}, " ")
}

type CreateLoadbalancerReqProvider struct {
	value string
}

type CreateLoadbalancerReqProviderEnum struct {
	VLB CreateLoadbalancerReqProvider
}

func GetCreateLoadbalancerReqProviderEnum() CreateLoadbalancerReqProviderEnum {
	return CreateLoadbalancerReqProviderEnum{
		VLB: CreateLoadbalancerReqProvider{
			value: "vlb",
		},
	}
}

func (c CreateLoadbalancerReqProvider) Value() string {
	return c.value
}

func (c CreateLoadbalancerReqProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadbalancerReqProvider) UnmarshalJSON(b []byte) error {
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

type CreateLoadbalancerReqProtectionStatus struct {
	value string
}

type CreateLoadbalancerReqProtectionStatusEnum struct {
	NON_PROTECTION     CreateLoadbalancerReqProtectionStatus
	CONSOLE_PROTECTION CreateLoadbalancerReqProtectionStatus
}

func GetCreateLoadbalancerReqProtectionStatusEnum() CreateLoadbalancerReqProtectionStatusEnum {
	return CreateLoadbalancerReqProtectionStatusEnum{
		NON_PROTECTION: CreateLoadbalancerReqProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateLoadbalancerReqProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateLoadbalancerReqProtectionStatus) Value() string {
	return c.value
}

func (c CreateLoadbalancerReqProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadbalancerReqProtectionStatus) UnmarshalJSON(b []byte) error {
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
