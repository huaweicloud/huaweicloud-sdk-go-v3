package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateCertificateResponse Response Object
type UpdateCertificateResponse struct {

	// SSL证书id
	Id *string `json:"id,omitempty"`

	// SSL证书所在的项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// SSL证书的管理状态；暂不支持
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// SSL证书的名称。
	Name *string `json:"name,omitempty"`

	// SSL证书的描述。
	Description *string `json:"description,omitempty"`

	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。
	Type *UpdateCertificateResponseType `json:"type,omitempty"`

	// 服务器证书所签域名。该字段仅type为server时有效。
	Domain *string `json:"domain,omitempty"`

	// 服务器证书的私钥。仅type为server时有效。type为server时必选。
	PrivateKey *string `json:"private_key,omitempty"`

	// 当type为server时，表示服务器证书的公钥；当type为client时，表示用于认证客户端证书的CA证书。
	Certificate *string `json:"certificate,omitempty"`

	// SSL证书的过期时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	ExpireTime *string `json:"expire_time,omitempty"`

	// SSL证书的创建时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	CreateTime *string `json:"create_time,omitempty"`

	// SSL证书的更新时间。 UTC时间，格式为：yyyy-MM-dd HH:mm:ss ，如2020-05-28 08:30:09
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCertificateResponse", string(data)}, " ")
}

type UpdateCertificateResponseType struct {
	value string
}

type UpdateCertificateResponseTypeEnum struct {
	SERVER UpdateCertificateResponseType
	CLIENT UpdateCertificateResponseType
}

func GetUpdateCertificateResponseTypeEnum() UpdateCertificateResponseTypeEnum {
	return UpdateCertificateResponseTypeEnum{
		SERVER: UpdateCertificateResponseType{
			value: "server",
		},
		CLIENT: UpdateCertificateResponseType{
			value: "client",
		},
	}
}

func (c UpdateCertificateResponseType) Value() string {
	return c.value
}

func (c UpdateCertificateResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCertificateResponseType) UnmarshalJSON(b []byte) error {
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
