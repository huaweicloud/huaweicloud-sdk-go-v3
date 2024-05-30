package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IdsParam struct {

	// ID列表，填写String类型替代Long类型。
	Ids []string `json:"ids"`

	// 删除物理表。 枚举值：   - PHYSICAL_TABLE: 关系建模
	DelTypes *IdsParamDelTypes `json:"del_types,omitempty"`
}

func (o IdsParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdsParam struct{}"
	}

	return strings.Join([]string{"IdsParam", string(data)}, " ")
}

type IdsParamDelTypes struct {
	value string
}

type IdsParamDelTypesEnum struct {
	PHYSICAL_TABLE IdsParamDelTypes
}

func GetIdsParamDelTypesEnum() IdsParamDelTypesEnum {
	return IdsParamDelTypesEnum{
		PHYSICAL_TABLE: IdsParamDelTypes{
			value: "PHYSICAL_TABLE",
		},
	}
}

func (c IdsParamDelTypes) Value() string {
	return c.value
}

func (c IdsParamDelTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IdsParamDelTypes) UnmarshalJSON(b []byte) error {
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
