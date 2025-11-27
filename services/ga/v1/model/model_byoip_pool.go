package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ByoipPool 自带IP地址池。
type ByoipPool struct {

	// 自带IP地址池ID。
	Id *string `json:"id,omitempty"`

	// 自带IP地址池的IP网段。
	Cidr *string `json:"cidr,omitempty"`

	// IP地址类型。 取值范围：IPV4、IPV6
	IpType *ByoipPoolIpType `json:"ip_type,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	Area *Area `json:"area,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`
}

func (o ByoipPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ByoipPool struct{}"
	}

	return strings.Join([]string{"ByoipPool", string(data)}, " ")
}

type ByoipPoolIpType struct {
	value string
}

type ByoipPoolIpTypeEnum struct {
	IPV4 ByoipPoolIpType
	IPV6 ByoipPoolIpType
}

func GetByoipPoolIpTypeEnum() ByoipPoolIpTypeEnum {
	return ByoipPoolIpTypeEnum{
		IPV4: ByoipPoolIpType{
			value: "IPV4",
		},
		IPV6: ByoipPoolIpType{
			value: "IPV6",
		},
	}
}

func (c ByoipPoolIpType) Value() string {
	return c.value
}

func (c ByoipPoolIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ByoipPoolIpType) UnmarshalJSON(b []byte) error {
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
