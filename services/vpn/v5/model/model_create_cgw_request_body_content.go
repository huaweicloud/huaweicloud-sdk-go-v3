package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateCgwRequestBodyContent struct {

	// 网关名称
	Name *string `json:"name,omitempty"`

	// 对端网关标识类型
	IdType *CreateCgwRequestBodyContentIdType `json:"id_type,omitempty"`

	// 对端网关标识值
	IdValue *string `json:"id_value,omitempty"`

	// 网关的bgp asn号，默认值为65000
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	CaCertificate *CaCertificateRequest `json:"ca_certificate,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`
}

func (o CreateCgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateCgwRequestBodyContent", string(data)}, " ")
}

type CreateCgwRequestBodyContentIdType struct {
	value string
}

type CreateCgwRequestBodyContentIdTypeEnum struct {
	IP   CreateCgwRequestBodyContentIdType
	FQDN CreateCgwRequestBodyContentIdType
}

func GetCreateCgwRequestBodyContentIdTypeEnum() CreateCgwRequestBodyContentIdTypeEnum {
	return CreateCgwRequestBodyContentIdTypeEnum{
		IP: CreateCgwRequestBodyContentIdType{
			value: "ip",
		},
		FQDN: CreateCgwRequestBodyContentIdType{
			value: "fqdn",
		},
	}
}

func (c CreateCgwRequestBodyContentIdType) Value() string {
	return c.value
}

func (c CreateCgwRequestBodyContentIdType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCgwRequestBodyContentIdType) UnmarshalJSON(b []byte) error {
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
