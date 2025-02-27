package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCertificatesRequest Request Object
type ListCertificatesRequest struct {

	// 每页返回的个数。 取值范围：0~intmax。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的证书id，为空时为查询第一页。 仅当和limit一起使用时生效
	Marker *string `json:"marker,omitempty"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 仅当和limit一起使用时生效。
	PageReverse *string `json:"page_reverse,omitempty"`

	// SSL证书ID。
	Id *string `json:"id,omitempty"`

	// SSL证书的名称。
	Name *string `json:"name,omitempty"`

	// 证书描述SSL证书描述。
	Description *string `json:"description,omitempty"`

	// SSL证书的类型。默认值：server；取值范围：server：服务端证书；client：客户端证书；
	Type *string `json:"type,omitempty"`

	// 服务端证书所签的域名。 取值：总长度为0-1024。  普通域名由若干字符串组成，字符串间以\".\"分割，单个字符串长度不超过63个字符，只能包含英文字母、数字或\"-\"，且必须以字母或数字开头和结尾。  泛域名仅允许首段为\"*\"，其他取值限制与普通域名一致。如：*.domain.com，但不能是：*my.domain.com   该字段仅type为server时有效。
	Domain *string `json:"domain,omitempty"`

	// PEM格式的服务端私有密钥。
	PrivateKey *string `json:"private_key,omitempty"`

	// PEM格式的服务端公有密钥或者用于认证客户端证书的CA证书，由type字段区分。
	Certificate *string `json:"certificate,omitempty"`

	// 参数解释： 证书来源  约束限制： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”。  取值范围： 无  默认取值： 当scm_certificate_id不为空，且未传入source时，默认取值为“scm”； 其他情况下默认为空。
	Source *string `json:"source,omitempty"`

	// 参数解释： 修改保护状态  约束限制： 无  取值范围：  - nonProtection: 不保护  - consoleProtection: 控制台修改保护  默认取值： nonProtection
	ProtectionStatus *ListCertificatesRequestProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释： 设置修改保护的原因  约束限制： 仅当protection_status为consoleProtection时有效  取值范围： 无  默认取值： 空
	ProtectionReason *string `json:"protection_reason,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}

type ListCertificatesRequestProtectionStatus struct {
	value string
}

type ListCertificatesRequestProtectionStatusEnum struct {
	NON_PROTECTION     ListCertificatesRequestProtectionStatus
	CONSOLE_PROTECTION ListCertificatesRequestProtectionStatus
}

func GetListCertificatesRequestProtectionStatusEnum() ListCertificatesRequestProtectionStatusEnum {
	return ListCertificatesRequestProtectionStatusEnum{
		NON_PROTECTION: ListCertificatesRequestProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: ListCertificatesRequestProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c ListCertificatesRequestProtectionStatus) Value() string {
	return c.value
}

func (c ListCertificatesRequestProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCertificatesRequestProtectionStatus) UnmarshalJSON(b []byte) error {
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
