package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAgencyRequest Request Object
type ShowAgencyRequest struct {

	// 委托类型,JOB_TRUST-任务委托,ADMIN_TRUST-系统委托
	AgencyType ShowAgencyRequestAgencyType `json:"agency_type"`
}

func (o ShowAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyRequest", string(data)}, " ")
}

type ShowAgencyRequestAgencyType struct {
	value string
}

type ShowAgencyRequestAgencyTypeEnum struct {
	JOB_TRUST   ShowAgencyRequestAgencyType
	ADMIN_TRUST ShowAgencyRequestAgencyType
}

func GetShowAgencyRequestAgencyTypeEnum() ShowAgencyRequestAgencyTypeEnum {
	return ShowAgencyRequestAgencyTypeEnum{
		JOB_TRUST: ShowAgencyRequestAgencyType{
			value: "JOB_TRUST",
		},
		ADMIN_TRUST: ShowAgencyRequestAgencyType{
			value: "ADMIN_TRUST",
		},
	}
}

func (c ShowAgencyRequestAgencyType) Value() string {
	return c.value
}

func (c ShowAgencyRequestAgencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAgencyRequestAgencyType) UnmarshalJSON(b []byte) error {
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
