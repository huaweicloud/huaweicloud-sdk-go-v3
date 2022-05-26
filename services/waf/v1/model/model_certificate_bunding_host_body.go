package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 绑定域名列表
type CertificateBundingHostBody struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`

	// waf模式（分为云模式：cloud,独享模式：premium）
	WafType *CertificateBundingHostBodyWafType `json:"waf_type,omitempty"`
}

func (o CertificateBundingHostBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateBundingHostBody struct{}"
	}

	return strings.Join([]string{"CertificateBundingHostBody", string(data)}, " ")
}

type CertificateBundingHostBodyWafType struct {
	value string
}

type CertificateBundingHostBodyWafTypeEnum struct {
	CLOUD   CertificateBundingHostBodyWafType
	PREMIUM CertificateBundingHostBodyWafType
}

func GetCertificateBundingHostBodyWafTypeEnum() CertificateBundingHostBodyWafTypeEnum {
	return CertificateBundingHostBodyWafTypeEnum{
		CLOUD: CertificateBundingHostBodyWafType{
			value: "cloud",
		},
		PREMIUM: CertificateBundingHostBodyWafType{
			value: "premium",
		},
	}
}

func (c CertificateBundingHostBodyWafType) Value() string {
	return c.value
}

func (c CertificateBundingHostBodyWafType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateBundingHostBodyWafType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
