package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceCheck struct {

	// 校验类型，仅支持校验创建实例。
	Action ResourceCheckAction `json:"action"`

	Resource *ResourceCheckResource `json:"resource"`
}

func (o ResourceCheck) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCheck struct{}"
	}

	return strings.Join([]string{"ResourceCheck", string(data)}, " ")
}

type ResourceCheckAction struct {
	value string
}

type ResourceCheckActionEnum struct {
	CREATE_INSTANCE ResourceCheckAction
}

func GetResourceCheckActionEnum() ResourceCheckActionEnum {
	return ResourceCheckActionEnum{
		CREATE_INSTANCE: ResourceCheckAction{
			value: "createInstance",
		},
	}
}

func (c ResourceCheckAction) Value() string {
	return c.value
}

func (c ResourceCheckAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceCheckAction) UnmarshalJSON(b []byte) error {
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
