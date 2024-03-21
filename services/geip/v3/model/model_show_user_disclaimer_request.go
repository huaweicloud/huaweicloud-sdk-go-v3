package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUserDisclaimerRequest Request Object
type ShowUserDisclaimerRequest struct {
	Fields *[]ShowUserDisclaimerRequestFields `json:"fields,omitempty"`
}

func (o ShowUserDisclaimerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserDisclaimerRequest struct{}"
	}

	return strings.Join([]string{"ShowUserDisclaimerRequest", string(data)}, " ")
}

type ShowUserDisclaimerRequestFields struct {
	value string
}

type ShowUserDisclaimerRequestFieldsEnum struct {
	DOMAIN_ID  ShowUserDisclaimerRequestFields
	CREATED_AT ShowUserDisclaimerRequestFields
	UPDATED_AT ShowUserDisclaimerRequestFields
}

func GetShowUserDisclaimerRequestFieldsEnum() ShowUserDisclaimerRequestFieldsEnum {
	return ShowUserDisclaimerRequestFieldsEnum{
		DOMAIN_ID: ShowUserDisclaimerRequestFields{
			value: "domain_id",
		},
		CREATED_AT: ShowUserDisclaimerRequestFields{
			value: "created_at",
		},
		UPDATED_AT: ShowUserDisclaimerRequestFields{
			value: "updated_at",
		},
	}
}

func (c ShowUserDisclaimerRequestFields) Value() string {
	return c.value
}

func (c ShowUserDisclaimerRequestFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserDisclaimerRequestFields) UnmarshalJSON(b []byte) error {
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
