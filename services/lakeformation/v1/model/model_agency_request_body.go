package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgencyRequestBody 委托请求信息
type AgencyRequestBody struct {

	// 委托类型，JOB_TRUST-任务委托,ADMIN_TRUST-系统委托
	AgencyType AgencyRequestBodyAgencyType `json:"agency_type"`
}

func (o AgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyRequestBody struct{}"
	}

	return strings.Join([]string{"AgencyRequestBody", string(data)}, " ")
}

type AgencyRequestBodyAgencyType struct {
	value string
}

type AgencyRequestBodyAgencyTypeEnum struct {
	JOB_TRUST   AgencyRequestBodyAgencyType
	ADMIN_TRUST AgencyRequestBodyAgencyType
}

func GetAgencyRequestBodyAgencyTypeEnum() AgencyRequestBodyAgencyTypeEnum {
	return AgencyRequestBodyAgencyTypeEnum{
		JOB_TRUST: AgencyRequestBodyAgencyType{
			value: "JOB_TRUST",
		},
		ADMIN_TRUST: AgencyRequestBodyAgencyType{
			value: "ADMIN_TRUST",
		},
	}
}

func (c AgencyRequestBodyAgencyType) Value() string {
	return c.value
}

func (c AgencyRequestBodyAgencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyRequestBodyAgencyType) UnmarshalJSON(b []byte) error {
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
