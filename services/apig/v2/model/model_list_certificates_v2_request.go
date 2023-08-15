package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCertificatesV2Request Request Object
type ListCertificatesV2Request struct {

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 证书名称
	Name *string `json:"name,omitempty"`

	// 证书域名
	CommonName *string `json:"common_name,omitempty"`

	// 证书签名算法
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 证书可见范围
	Type *ListCertificatesV2RequestType `json:"type,omitempty"`

	// 证书所属实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListCertificatesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesV2Request struct{}"
	}

	return strings.Join([]string{"ListCertificatesV2Request", string(data)}, " ")
}

type ListCertificatesV2RequestType struct {
	value string
}

type ListCertificatesV2RequestTypeEnum struct {
	INSTANCE ListCertificatesV2RequestType
	GLOBAL   ListCertificatesV2RequestType
}

func GetListCertificatesV2RequestTypeEnum() ListCertificatesV2RequestTypeEnum {
	return ListCertificatesV2RequestTypeEnum{
		INSTANCE: ListCertificatesV2RequestType{
			value: "instance",
		},
		GLOBAL: ListCertificatesV2RequestType{
			value: "global",
		},
	}
}

func (c ListCertificatesV2RequestType) Value() string {
	return c.value
}

func (c ListCertificatesV2RequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCertificatesV2RequestType) UnmarshalJSON(b []byte) error {
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
