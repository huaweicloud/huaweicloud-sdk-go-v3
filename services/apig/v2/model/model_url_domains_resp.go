package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UrlDomainsResp struct {
	// 域名编号

	Id *string `json:"id,omitempty"`
	// 访问域名

	Domain *string `json:"domain,omitempty"`
	// 域名cname状态： - 1：未解析 - 2：解析中 - 3：解析成功 - 4：解析失败

	CnameStatus *int32 `json:"cname_status,omitempty"`
	// SSL证书编号

	SslId *string `json:"ssl_id,omitempty"`
	// SSL证书名称

	SslName *string `json:"ssl_name,omitempty"`
	// 最小SSL协议版本号，支持TLSv1.1或TLSv1.2

	MinSslVersion *UrlDomainsRespMinSslVersion `json:"min_ssl_version,omitempty"`
}

func (o UrlDomainsResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UrlDomainsResp struct{}"
	}

	return strings.Join([]string{"UrlDomainsResp", string(data)}, " ")
}

type UrlDomainsRespMinSslVersion struct {
	value string
}

type UrlDomainsRespMinSslVersionEnum struct {
	TL_SV1_1 UrlDomainsRespMinSslVersion
	TL_SV1_2 UrlDomainsRespMinSslVersion
}

func GetUrlDomainsRespMinSslVersionEnum() UrlDomainsRespMinSslVersionEnum {
	return UrlDomainsRespMinSslVersionEnum{
		TL_SV1_1: UrlDomainsRespMinSslVersion{
			value: "TLSv1.1",
		},
		TL_SV1_2: UrlDomainsRespMinSslVersion{
			value: "TLSv1.2",
		},
	}
}

func (c UrlDomainsRespMinSslVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UrlDomainsRespMinSslVersion) UnmarshalJSON(b []byte) error {
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
