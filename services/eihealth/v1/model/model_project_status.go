package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProjectStatus struct {
	value string
}

type ProjectStatusEnum struct {
	ACTIVE   ProjectStatus
	INACTIVE ProjectStatus
}

func GetProjectStatusEnum() ProjectStatusEnum {
	return ProjectStatusEnum{
		ACTIVE: ProjectStatus{
			value: "ACTIVE",
		},
		INACTIVE: ProjectStatus{
			value: "INACTIVE",
		},
	}
}

func (c ProjectStatus) Value() string {
	return c.value
}

func (c ProjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectStatus) UnmarshalJSON(b []byte) error {
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
