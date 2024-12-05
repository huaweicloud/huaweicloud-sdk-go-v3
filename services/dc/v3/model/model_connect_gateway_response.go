package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ConnectGatewayResponse 互联网关的相信信息对象
type ConnectGatewayResponse struct {

	// 唯一ID
	Id *string `json:"id,omitempty"`

	// 租户项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 网关名字
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 地址族信息 - ipv4: 仅支持ipv4模式 - dual: 支持ipv4 和 ipv6 模式
	AddressFamily *ConnectGatewayResponseAddressFamily `json:"address_family,omitempty"`

	// 网关状态 - DOWN 未使用或关联设备状态为DOWN - ACTIVE 正常 - ERROR 异常
	Status *string `json:"status,omitempty"`

	// 网关站点值
	AccessSite *string `json:"access_site,omitempty"`

	// BGP类型AS号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 当前绑定的global eip数量
	CurrentGeipCount *int32 `json:"current_geip_count,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 带宽包id
	GcbId *string `json:"gcb_id,omitempty"`

	// 网关位置
	GatewaySite *string `json:"gateway_site,omitempty"`
}

func (o ConnectGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectGatewayResponse struct{}"
	}

	return strings.Join([]string{"ConnectGatewayResponse", string(data)}, " ")
}

type ConnectGatewayResponseAddressFamily struct {
	value string
}

type ConnectGatewayResponseAddressFamilyEnum struct {
	IPV4 ConnectGatewayResponseAddressFamily
	DUAL ConnectGatewayResponseAddressFamily
}

func GetConnectGatewayResponseAddressFamilyEnum() ConnectGatewayResponseAddressFamilyEnum {
	return ConnectGatewayResponseAddressFamilyEnum{
		IPV4: ConnectGatewayResponseAddressFamily{
			value: "ipv4",
		},
		DUAL: ConnectGatewayResponseAddressFamily{
			value: "dual",
		},
	}
}

func (c ConnectGatewayResponseAddressFamily) Value() string {
	return c.value
}

func (c ConnectGatewayResponseAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectGatewayResponseAddressFamily) UnmarshalJSON(b []byte) error {
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
