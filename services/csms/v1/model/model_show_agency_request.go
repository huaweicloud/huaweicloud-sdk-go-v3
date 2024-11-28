package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAgencyRequest Request Object
type ShowAgencyRequest struct {

	// 凭据类型。
	SecretType ShowAgencyRequestSecretType `json:"secret_type"`
}

func (o ShowAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyRequest", string(data)}, " ")
}

type ShowAgencyRequestSecretType struct {
	value string
}

type ShowAgencyRequestSecretTypeEnum struct {
	RDS_FG      ShowAgencyRequestSecretType
	GAUSS_DB_FG ShowAgencyRequestSecretType
}

func GetShowAgencyRequestSecretTypeEnum() ShowAgencyRequestSecretTypeEnum {
	return ShowAgencyRequestSecretTypeEnum{
		RDS_FG: ShowAgencyRequestSecretType{
			value: "RDS-FG",
		},
		GAUSS_DB_FG: ShowAgencyRequestSecretType{
			value: "GaussDB-FG",
		},
	}
}

func (c ShowAgencyRequestSecretType) Value() string {
	return c.value
}

func (c ShowAgencyRequestSecretType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyRequestSecretType) UnmarshalJSON(b []byte) error {
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
