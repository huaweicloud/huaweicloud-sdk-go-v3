package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateGlobalDcGateway 创建global-dc-gateway的属性详情。
type CreateGlobalDcGateway struct {

	// 项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// global-dc-gateway名字。
	Name string `json:"name"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// global-dc-gateway对应的ASN号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// global-dc-gateway所属的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 网关的地址簇，IPv4或者ipv6和IPv4双栈 - ipv4 - dual
	AddressFamily *CreateGlobalDcGatewayAddressFamily `json:"address_family,omitempty"`

	// global-dc-gateway关联TAG。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateGlobalDcGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGateway struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGateway", string(data)}, " ")
}

type CreateGlobalDcGatewayAddressFamily struct {
	value string
}

type CreateGlobalDcGatewayAddressFamilyEnum struct {
	IPV4 CreateGlobalDcGatewayAddressFamily
	DUAL CreateGlobalDcGatewayAddressFamily
}

func GetCreateGlobalDcGatewayAddressFamilyEnum() CreateGlobalDcGatewayAddressFamilyEnum {
	return CreateGlobalDcGatewayAddressFamilyEnum{
		IPV4: CreateGlobalDcGatewayAddressFamily{
			value: "ipv4",
		},
		DUAL: CreateGlobalDcGatewayAddressFamily{
			value: "dual",
		},
	}
}

func (c CreateGlobalDcGatewayAddressFamily) Value() string {
	return c.value
}

func (c CreateGlobalDcGatewayAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGlobalDcGatewayAddressFamily) UnmarshalJSON(b []byte) error {
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
