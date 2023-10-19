package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAgencyResponse Response Object
type ShowAgencyResponse struct {

	// 委托类型,JOB_TRUST-任务委托,ADMIN_TRUST-系统委托
	AgencyType     *ShowAgencyResponseAgencyType `json:"agency_type,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}

type ShowAgencyResponseAgencyType struct {
	value string
}

type ShowAgencyResponseAgencyTypeEnum struct {
	JOB_TRUST   ShowAgencyResponseAgencyType
	ADMIN_TRUST ShowAgencyResponseAgencyType
}

func GetShowAgencyResponseAgencyTypeEnum() ShowAgencyResponseAgencyTypeEnum {
	return ShowAgencyResponseAgencyTypeEnum{
		JOB_TRUST: ShowAgencyResponseAgencyType{
			value: "JOB_TRUST",
		},
		ADMIN_TRUST: ShowAgencyResponseAgencyType{
			value: "ADMIN_TRUST",
		},
	}
}

func (c ShowAgencyResponseAgencyType) Value() string {
	return c.value
}

func (c ShowAgencyResponseAgencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyResponseAgencyType) UnmarshalJSON(b []byte) error {
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
