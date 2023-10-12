package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertForm 添加或编辑证书的请求体表单
type CertForm struct {

	// 证书名称。长度为4 ~ 50位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。
	Name string `json:"name"`

	// 证书内容
	CertContent string `json:"cert_content"`

	// 证书私钥
	PrivateKey string `json:"private_key"`

	// 证书可见范围
	Type *CertFormType `json:"type,omitempty"`

	// 所属实例ID，当type=instance时必填
	InstanceId *string `json:"instance_id,omitempty"`

	// 信任的根证书CA  [暂不支持](tag:fcs;hcs;g42;Site)
	TrustedRootCa *string `json:"trusted_root_ca,omitempty"`
}

func (o CertForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertForm struct{}"
	}

	return strings.Join([]string{"CertForm", string(data)}, " ")
}

type CertFormType struct {
	value string
}

type CertFormTypeEnum struct {
	INSTANCE CertFormType
	GLOBAL   CertFormType
}

func GetCertFormTypeEnum() CertFormTypeEnum {
	return CertFormTypeEnum{
		INSTANCE: CertFormType{
			value: "instance",
		},
		GLOBAL: CertFormType{
			value: "global",
		},
	}
}

func (c CertFormType) Value() string {
	return c.value
}

func (c CertFormType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertFormType) UnmarshalJSON(b []byte) error {
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
