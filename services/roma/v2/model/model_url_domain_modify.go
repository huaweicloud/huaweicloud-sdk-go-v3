package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UrlDomainModify struct {

	// 最小ssl协议版本号。支持TLSv1.1或TLSv1.2
	MinSslVersion *UrlDomainModifyMinSslVersion `json:"min_ssl_version,omitempty"`
}

func (o UrlDomainModify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDomainModify struct{}"
	}

	return strings.Join([]string{"UrlDomainModify", string(data)}, " ")
}

type UrlDomainModifyMinSslVersion struct {
	value string
}

type UrlDomainModifyMinSslVersionEnum struct {
	TL_SV1_1 UrlDomainModifyMinSslVersion
	TL_SV1_2 UrlDomainModifyMinSslVersion
}

func GetUrlDomainModifyMinSslVersionEnum() UrlDomainModifyMinSslVersionEnum {
	return UrlDomainModifyMinSslVersionEnum{
		TL_SV1_1: UrlDomainModifyMinSslVersion{
			value: "TLSv1.1",
		},
		TL_SV1_2: UrlDomainModifyMinSslVersion{
			value: "TLSv1.2",
		},
	}
}

func (c UrlDomainModifyMinSslVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainModifyMinSslVersion) UnmarshalJSON(b []byte) error {
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
