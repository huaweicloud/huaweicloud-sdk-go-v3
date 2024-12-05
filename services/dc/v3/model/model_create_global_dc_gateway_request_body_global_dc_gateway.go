package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateGlobalDcGatewayRequestBodyGlobalDcGateway struct {

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 全球接入网关名称
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// BGP模式的AS号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 地址簇类型
	AddressFamily *CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily `json:"address_family,omitempty"`

	// 标签
	Tags *[]CreateGlobalDcGatewayRequestBodyGlobalDcGatewayTags `json:"tags,omitempty"`
}

func (o CreateGlobalDcGatewayRequestBodyGlobalDcGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayRequestBodyGlobalDcGateway struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayRequestBodyGlobalDcGateway", string(data)}, " ")
}

type CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily struct {
	value string
}

type CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum struct {
	IPV4 CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily
	DUAL CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily
}

func GetCreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum() CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum {
	return CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamilyEnum{
		IPV4: CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily{
			value: "ipv4",
		},
		DUAL: CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily{
			value: "dual",
		},
	}
}

func (c CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily) Value() string {
	return c.value
}

func (c CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGlobalDcGatewayRequestBodyGlobalDcGatewayAddressFamily) UnmarshalJSON(b []byte) error {
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
