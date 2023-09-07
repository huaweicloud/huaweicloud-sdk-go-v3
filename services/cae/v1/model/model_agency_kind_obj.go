package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgencyKindObj API类型，固定值“Agency”，该值不可修改。
type AgencyKindObj struct {
	value string
}

type AgencyKindObjEnum struct {
	AGENCY AgencyKindObj
}

func GetAgencyKindObjEnum() AgencyKindObjEnum {
	return AgencyKindObjEnum{
		AGENCY: AgencyKindObj{
			value: "Agency",
		},
	}
}

func (c AgencyKindObj) Value() string {
	return c.value
}

func (c AgencyKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyKindObj) UnmarshalJSON(b []byte) error {
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
