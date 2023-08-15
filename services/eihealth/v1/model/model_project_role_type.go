package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProjectRoleType struct {
	value string
}

type ProjectRoleTypeEnum struct {
	ADMINISTRATOR ProjectRoleType
	DEVELOPER     ProjectRoleType
	UPLOADER      ProjectRoleType
	VIEWER        ProjectRoleType
}

func GetProjectRoleTypeEnum() ProjectRoleTypeEnum {
	return ProjectRoleTypeEnum{
		ADMINISTRATOR: ProjectRoleType{
			value: "Administrator",
		},
		DEVELOPER: ProjectRoleType{
			value: "Developer",
		},
		UPLOADER: ProjectRoleType{
			value: "Uploader",
		},
		VIEWER: ProjectRoleType{
			value: "Viewer",
		},
	}
}

func (c ProjectRoleType) Value() string {
	return c.value
}

func (c ProjectRoleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectRoleType) UnmarshalJSON(b []byte) error {
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
