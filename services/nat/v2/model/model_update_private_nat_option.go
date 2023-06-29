package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePrivateNatOption 更新私网NAT网关实例的请求体。
type UpdatePrivateNatOption struct {

	// 私网NAT网关实例的名字。 私网NAT网关实例的名字仅支持数字、字母、_（下划线）、-（中划线）、中文。
	Name *string `json:"name,omitempty"`

	// 私网NAT网关的描述。
	Description *string `json:"description,omitempty"`

	// 私网NAT网关实例的规格。 取值为： \"Small\"：小型 \"Medium\"：中型 \"Large\"：大型 \"Extra-large\"：超大型
	Spec *UpdatePrivateNatOptionSpec `json:"spec,omitempty"`
}

func (o UpdatePrivateNatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateNatOption struct{}"
	}

	return strings.Join([]string{"UpdatePrivateNatOption", string(data)}, " ")
}

type UpdatePrivateNatOptionSpec struct {
	value string
}

type UpdatePrivateNatOptionSpecEnum struct {
	SMALL       UpdatePrivateNatOptionSpec
	MEDIUM      UpdatePrivateNatOptionSpec
	LARGE       UpdatePrivateNatOptionSpec
	EXTRA_LARGE UpdatePrivateNatOptionSpec
}

func GetUpdatePrivateNatOptionSpecEnum() UpdatePrivateNatOptionSpecEnum {
	return UpdatePrivateNatOptionSpecEnum{
		SMALL: UpdatePrivateNatOptionSpec{
			value: "Small",
		},
		MEDIUM: UpdatePrivateNatOptionSpec{
			value: "Medium",
		},
		LARGE: UpdatePrivateNatOptionSpec{
			value: "Large",
		},
		EXTRA_LARGE: UpdatePrivateNatOptionSpec{
			value: "Extra-large",
		},
	}
}

func (c UpdatePrivateNatOptionSpec) Value() string {
	return c.value
}

func (c UpdatePrivateNatOptionSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePrivateNatOptionSpec) UnmarshalJSON(b []byte) error {
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
