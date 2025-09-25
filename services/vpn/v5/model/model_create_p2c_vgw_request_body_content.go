package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateP2cVgwRequestBodyContent struct {

	// P2C VPN网关名称。1-64字符，中文、英文、数字包含下划线
	Name *string `json:"name,omitempty"`

	// P2C VPN网关所连接的VPC的ID。标准UUID
	VpcId string `json:"vpc_id"`

	// P2C VPN网关所使用的VPC子网ID。标准的UUID
	ConnectSubnet string `json:"connect_subnet"`

	// P2C VPN网关的规格类型
	Flavor *CreateP2cVgwRequestBodyContentFlavor `json:"flavor,omitempty"`

	// 不填时自动为P2C VPN网关选择可用区。如果需要指定可用区可以通过查询VPN网关可用区查询可用区列表
	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty"`

	Eip *CreateP2cRequestEip `json:"eip"`

	// 最大并发客户端连接数，且不超过规格的最大连接数
	MaxConnectionNumber *int32 `json:"max_connection_number,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o CreateP2cVgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2cVgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateP2cVgwRequestBodyContent", string(data)}, " ")
}

type CreateP2cVgwRequestBodyContentFlavor struct {
	value string
}

type CreateP2cVgwRequestBodyContentFlavorEnum struct {
	PROFESSIONAL1 CreateP2cVgwRequestBodyContentFlavor
}

func GetCreateP2cVgwRequestBodyContentFlavorEnum() CreateP2cVgwRequestBodyContentFlavorEnum {
	return CreateP2cVgwRequestBodyContentFlavorEnum{
		PROFESSIONAL1: CreateP2cVgwRequestBodyContentFlavor{
			value: "Professional1",
		},
	}
}

func (c CreateP2cVgwRequestBodyContentFlavor) Value() string {
	return c.value
}

func (c CreateP2cVgwRequestBodyContentFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateP2cVgwRequestBodyContentFlavor) UnmarshalJSON(b []byte) error {
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
