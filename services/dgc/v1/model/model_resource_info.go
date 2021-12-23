package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceInfo struct {
	Name *string `json:"name,omitempty"`

	Type *ResourceInfoType `json:"type,omitempty"`
	// 资源文件所在OBS路径

	Location *string `json:"location,omitempty"`

	DependFiles *[]string `json:"dependFiles,omitempty"`

	Desc *string `json:"desc,omitempty"`
	// 资源所在目录

	Directory *string `json:"directory,omitempty"`
}

func (o ResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInfo struct{}"
	}

	return strings.Join([]string{"ResourceInfo", string(data)}, " ")
}

type ResourceInfoType struct {
	value string
}

type ResourceInfoTypeEnum struct {
	ARCHIVE ResourceInfoType
	FILE    ResourceInfoType
	JAR     ResourceInfoType
}

func GetResourceInfoTypeEnum() ResourceInfoTypeEnum {
	return ResourceInfoTypeEnum{
		ARCHIVE: ResourceInfoType{
			value: "archive",
		},
		FILE: ResourceInfoType{
			value: "file",
		},
		JAR: ResourceInfoType{
			value: "jar",
		},
	}
}

func (c ResourceInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceInfoType) UnmarshalJSON(b []byte) error {
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
