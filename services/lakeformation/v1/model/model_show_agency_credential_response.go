package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAgencyCredentialResponse Response Object
type ShowAgencyCredentialResponse struct {

	// 委托类型：TABLE_SERVICE_TRUST-表服务委托。
	AgencyType *ShowAgencyCredentialResponseAgencyType `json:"agency_type,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// security token
	SecurityToken *string `json:"security_token,omitempty"`

	// token到期时间
	ExpiresAt *string `json:"expires_at,omitempty"`

	// token下发时间
	IssuedAt *string `json:"issued_at,omitempty"`

	// ak
	Access *string `json:"access,omitempty"`

	// sk
	Secret         *string `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgencyCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyCredentialResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyCredentialResponse", string(data)}, " ")
}

type ShowAgencyCredentialResponseAgencyType struct {
	value string
}

type ShowAgencyCredentialResponseAgencyTypeEnum struct {
	TABLE_SERVICE_TRUST ShowAgencyCredentialResponseAgencyType
}

func GetShowAgencyCredentialResponseAgencyTypeEnum() ShowAgencyCredentialResponseAgencyTypeEnum {
	return ShowAgencyCredentialResponseAgencyTypeEnum{
		TABLE_SERVICE_TRUST: ShowAgencyCredentialResponseAgencyType{
			value: "TABLE_SERVICE_TRUST",
		},
	}
}

func (c ShowAgencyCredentialResponseAgencyType) Value() string {
	return c.value
}

func (c ShowAgencyCredentialResponseAgencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyCredentialResponseAgencyType) UnmarshalJSON(b []byte) error {
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
