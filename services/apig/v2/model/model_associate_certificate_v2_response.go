package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type AssociateCertificateV2Response struct {

	// 自定义域名
	UrlDomain string `json:"url_domain"`

	// 自定义域名的编号
	Id string `json:"id"`

	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败
	Status AssociateCertificateV2ResponseStatus `json:"status"`

	// 支持的最小SSL版本
	MinSslVersion string `json:"min_ssl_version"`

	// 证书的名称
	SslName string `json:"ssl_name"`

	// 证书的编号
	SslId          string `json:"ssl_id"`
	HttpStatusCode int    `json:"-"`
}

func (o AssociateCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateCertificateV2Response struct{}"
	}

	return strings.Join([]string{"AssociateCertificateV2Response", string(data)}, " ")
}

type AssociateCertificateV2ResponseStatus struct {
	value int32
}

type AssociateCertificateV2ResponseStatusEnum struct {
	E_1 AssociateCertificateV2ResponseStatus
	E_2 AssociateCertificateV2ResponseStatus
	E_3 AssociateCertificateV2ResponseStatus
	E_4 AssociateCertificateV2ResponseStatus
}

func GetAssociateCertificateV2ResponseStatusEnum() AssociateCertificateV2ResponseStatusEnum {
	return AssociateCertificateV2ResponseStatusEnum{
		E_1: AssociateCertificateV2ResponseStatus{
			value: 1,
		}, E_2: AssociateCertificateV2ResponseStatus{
			value: 2,
		}, E_3: AssociateCertificateV2ResponseStatus{
			value: 3,
		}, E_4: AssociateCertificateV2ResponseStatus{
			value: 4,
		},
	}
}

func (c AssociateCertificateV2ResponseStatus) Value() int32 {
	return c.value
}

func (c AssociateCertificateV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateCertificateV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
