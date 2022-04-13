package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建证书请求参数。
type CreateCertificateOption struct {
	// 证书的管理状态。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// HTTPS协议使用的证书内容。 如果type为server_sm时，该字段填写SM签名证书内容。 取值范围：PEM编码格式。

	Certificate string `json:"certificate"`
	// 证书的描述。

	Description *string `json:"description,omitempty"`
	// 服务器证书所签域名。该字段仅type为server或server_sm时有效。  总长度为0-1024，由若干普通域名或泛域名组成，域名之间以\",\"分割，不超过30个域名。  普通域名：由若干字符串组成，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。例：www.test.com；  泛域名：在普通域名的基础上仅允许首字母为\"\"。例：.test.com

	Domain *string `json:"domain,omitempty"`
	// 证书的名称。

	Name *string `json:"name,omitempty"`
	// HTTPS协议使用的私钥。仅type为server或server_sm时有效。type为server或server_sm时必选。 如果type为server_sm时，该字段填写SM签名证书的私钥。 取值范围：PEM编码格式。

	PrivateKey *string `json:"private_key,omitempty"`
	// 证书所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// SSL证书的类型。分为服务器证书(server)、CA证书(client)和服务器SM双证书(server_sm)。 默认值：server

	Type *CreateCertificateOptionType `json:"type,omitempty"`
	// 证书所属的企业项目ID。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// HTTPS协议使用的SM加密证书内容。  取值：PEM编码格式。  使用说明： - 仅type为server_sm时有效且必选。

	EncCertificate *string `json:"enc_certificate,omitempty"`
	// HTTPS协议使用的SM加密证书私钥。  取值：PEM编码格式。  使用说明： - 仅type为server_sm时有效且必选。

	EncPrivateKey *string `json:"enc_private_key,omitempty"`
}

func (o CreateCertificateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateOption struct{}"
	}

	return strings.Join([]string{"CreateCertificateOption", string(data)}, " ")
}

type CreateCertificateOptionType struct {
	value string
}

type CreateCertificateOptionTypeEnum struct {
	SERVER CreateCertificateOptionType
	CLIENT CreateCertificateOptionType
}

func GetCreateCertificateOptionTypeEnum() CreateCertificateOptionTypeEnum {
	return CreateCertificateOptionTypeEnum{
		SERVER: CreateCertificateOptionType{
			value: "server",
		},
		CLIENT: CreateCertificateOptionType{
			value: "client",
		},
	}
}

func (c CreateCertificateOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCertificateOptionType) UnmarshalJSON(b []byte) error {
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
