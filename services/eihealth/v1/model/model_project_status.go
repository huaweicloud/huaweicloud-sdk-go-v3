package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
