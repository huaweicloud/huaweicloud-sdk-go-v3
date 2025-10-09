package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Principal struct {

	// **参数解释：** 授权对象的类型。 **约束限制：** 不涉及 **取值范围：** - user：用户 - group：用户组 - agency：委托账号  **默认取值：** 不涉及
	Type PrincipalType `json:"type"`

	// **参数解释：** 授权对象ID列表，根据授权对象的类型，用户、用户组和委托账号，填入对应的ID。 **约束限制：** 当前最多支持同时授权500个用户/用户组。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ids []string `json:"ids"`
}

func (o Principal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Principal struct{}"
	}

	return strings.Join([]string{"Principal", string(data)}, " ")
}

type PrincipalType struct {
	value string
}

type PrincipalTypeEnum struct {
	USER   PrincipalType
	GROUP  PrincipalType
	AGENCY PrincipalType
}

func GetPrincipalTypeEnum() PrincipalTypeEnum {
	return PrincipalTypeEnum{
		USER: PrincipalType{
			value: "user",
		},
		GROUP: PrincipalType{
			value: "group",
		},
		AGENCY: PrincipalType{
			value: "agency",
		},
	}
}

func (c PrincipalType) Value() string {
	return c.value
}

func (c PrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrincipalType) UnmarshalJSON(b []byte) error {
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
