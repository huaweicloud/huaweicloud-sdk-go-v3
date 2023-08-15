package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// NeutronNetwork network对象列表
type NeutronNetwork struct {

	// 网络ID
	Id string `json:"id"`

	// 功能说明：网络状态 取值范围：ACTIVE，DOWN，BUILD或ERROR
	Status NeutronNetworkStatus `json:"status"`

	// 功能说明：网络名称 取值范围：0-255个字符
	Name string `json:"name"`

	// 功能说明：网络关联的子网ID列表 约束：一个network仅支持关联一个 subnet。
	Subnets []string `json:"subnets"`

	// 功能说明：扩展属性，是否外部网络 取值范围：true、false 约束：不支持设置和更新
	Routerexternal bool `json:"router:external"`

	// 功能说明：资源的管理状态 取值范围：true、false 约束：只支持true
	AdminStateUp bool `json:"admin_state_up"`

	// 功能说明：是否支持跨租户共享此资源 取值范围：true(共享)、false(非共享) 约束：不支持设置和更新
	Shared bool `json:"shared"`

	// 项目ID
	TenantId string `json:"tenant_id"`

	// 功能说明：扩展属性，网络类型（支持vxlan, geneve） 取值范围：vxlan，geneve
	ProvidernetworkType string `json:"provider:network_type"`

	// 功能说明：本网络的候选可用域
	AvailabilityZoneHints []string `json:"availability_zone_hints"`

	// 功能说明：本网络的可用域。 取值范围：当前region下的可用域
	AvailabilityZones []string `json:"availability_zones"`

	// 功能说明：端口安全使能标记 取值范围：true（启用）、false（禁用） 约束：如果不使能，则network下所有虚机的安全组和dhcp防欺骗不生效
	PortSecurityEnabled bool `json:"port_security_enabled"`

	// 功能说明：默认内网DNS域地址 约束：系统自动生成维护，不支持设置和更新
	DnsDomain string `json:"dns_domain"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NeutronNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronNetwork struct{}"
	}

	return strings.Join([]string{"NeutronNetwork", string(data)}, " ")
}

type NeutronNetworkStatus struct {
	value string
}

type NeutronNetworkStatusEnum struct {
	ACTIVE NeutronNetworkStatus
	DOWN   NeutronNetworkStatus
	BUILD  NeutronNetworkStatus
	ERROR  NeutronNetworkStatus
}

func GetNeutronNetworkStatusEnum() NeutronNetworkStatusEnum {
	return NeutronNetworkStatusEnum{
		ACTIVE: NeutronNetworkStatus{
			value: "ACTIVE",
		},
		DOWN: NeutronNetworkStatus{
			value: "DOWN",
		},
		BUILD: NeutronNetworkStatus{
			value: "BUILD",
		},
		ERROR: NeutronNetworkStatus{
			value: "ERROR",
		},
	}
}

func (c NeutronNetworkStatus) Value() string {
	return c.value
}

func (c NeutronNetworkStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronNetworkStatus) UnmarshalJSON(b []byte) error {
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
