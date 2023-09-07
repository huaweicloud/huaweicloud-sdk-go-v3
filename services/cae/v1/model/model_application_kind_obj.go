package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationKindObj API类型，固定值“Application”，该值不可修改。
type ApplicationKindObj struct {
	value string
}

type ApplicationKindObjEnum struct {
	APPLICATION ApplicationKindObj
}

func GetApplicationKindObjEnum() ApplicationKindObjEnum {
	return ApplicationKindObjEnum{
		APPLICATION: ApplicationKindObj{
			value: "Application",
		},
	}
}

func (c ApplicationKindObj) Value() string {
	return c.value
}

func (c ApplicationKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationKindObj) UnmarshalJSON(b []byte) error {
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
