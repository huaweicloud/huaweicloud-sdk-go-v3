package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AgencyItem struct {

	// 委托名称，固定值“cae_trust”，该值不可修改。
	Name *AgencyItemName `json:"name,omitempty"`
}

func (o AgencyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyItem struct{}"
	}

	return strings.Join([]string{"AgencyItem", string(data)}, " ")
}

type AgencyItemName struct {
	value string
}

type AgencyItemNameEnum struct {
	CAE_TRUST AgencyItemName
}

func GetAgencyItemNameEnum() AgencyItemNameEnum {
	return AgencyItemNameEnum{
		CAE_TRUST: AgencyItemName{
			value: "cae_trust",
		},
	}
}

func (c AgencyItemName) Value() string {
	return c.value
}

func (c AgencyItemName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgencyItemName) UnmarshalJSON(b []byte) error {
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
