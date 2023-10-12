package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCertificatesV2Request Request Object
type ListCertificatesV2Request struct {

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
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

	// 证书算法类型： - RSA - ECC - SM2  [暂不支持](tag:hws;hws_hk;g42;Site)
	AlgorithmType *ListCertificatesV2RequestAlgorithmType `json:"algorithm_type,omitempty"`
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

type ListCertificatesV2RequestAlgorithmType struct {
	value string
}

type ListCertificatesV2RequestAlgorithmTypeEnum struct {
	RSA ListCertificatesV2RequestAlgorithmType
	ECC ListCertificatesV2RequestAlgorithmType
	SM2 ListCertificatesV2RequestAlgorithmType
}

func GetListCertificatesV2RequestAlgorithmTypeEnum() ListCertificatesV2RequestAlgorithmTypeEnum {
	return ListCertificatesV2RequestAlgorithmTypeEnum{
		RSA: ListCertificatesV2RequestAlgorithmType{
			value: "RSA",
		},
		ECC: ListCertificatesV2RequestAlgorithmType{
			value: "ECC",
		},
		SM2: ListCertificatesV2RequestAlgorithmType{
			value: "SM2",
		},
	}
}

func (c ListCertificatesV2RequestAlgorithmType) Value() string {
	return c.value
}

func (c ListCertificatesV2RequestAlgorithmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCertificatesV2RequestAlgorithmType) UnmarshalJSON(b []byte) error {
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
