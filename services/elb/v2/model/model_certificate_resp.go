package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 证书接口增删改查响应体
type CertificateResp struct {
	// SSL证书id

	Id string `json:"id"`
	// SSL证书所在的项目ID

	TenantId string `json:"tenant_id"`
	// SSL证书的管理状态；暂不支持

	AdminStateUp bool `json:"admin_state_up"`
	// SSL证书的名称。

	Name string `json:"name"`
	// SSL证书的描述。

	Description string `json:"description"`
	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。

	Type CertificateRespType `json:"type"`
	// 服务器证书所签域名。该字段仅type为server时有效。

	Domain string `json:"domain"`
	// 服务器证书的私钥。仅type为server时有效。type为server时必选。

	PrivateKey string `json:"private_key"`
	// 当type为server时，表示服务器证书的公钥；当type为client时，表示用于认证客户端证书的CA证书。

	Certificate string `json:"certificate"`
	// SSL证书的过期时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09

	ExpireTime string `json:"expire_time"`
	// SSL证书的创建时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09

	CreateTime string `json:"create_time"`
	// SSL证书的更新时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09

	UpdateTime string `json:"update_time"`
}

func (o CertificateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateResp struct{}"
	}

	return strings.Join([]string{"CertificateResp", string(data)}, " ")
}

type CertificateRespType struct {
	value string
}

type CertificateRespTypeEnum struct {
	SERVER CertificateRespType
	CLIENT CertificateRespType
}

func GetCertificateRespTypeEnum() CertificateRespTypeEnum {
	return CertificateRespTypeEnum{
		SERVER: CertificateRespType{
			value: "server",
		},
		CLIENT: CertificateRespType{
			value: "client",
		},
	}
}

func (c CertificateRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertificateRespType) UnmarshalJSON(b []byte) error {
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
