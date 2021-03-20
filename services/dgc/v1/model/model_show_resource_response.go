package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowResourceResponse struct {
	Name *string `json:"name,omitempty"`

	Type *ShowResourceResponseType `json:"type,omitempty"`
	// 资源文件所在OBS路径

	Location *string `json:"location,omitempty"`

	DependFiles *[]string `json:"dependFiles,omitempty"`

	Desc *string `json:"desc,omitempty"`
	// 资源所在目录

	Directory      *string `json:"directory,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceResponse", string(data)}, " ")
}

type ShowResourceResponseType struct {
	value string
}

type ShowResourceResponseTypeEnum struct {
	ARCHIVE ShowResourceResponseType
	FILE    ShowResourceResponseType
	JAR     ShowResourceResponseType
}

func GetShowResourceResponseTypeEnum() ShowResourceResponseTypeEnum {
	return ShowResourceResponseTypeEnum{
		ARCHIVE: ShowResourceResponseType{
			value: "archive",
		},
		FILE: ShowResourceResponseType{
			value: "file",
		},
		JAR: ShowResourceResponseType{
			value: "jar",
		},
	}
}

func (c ShowResourceResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowResourceResponseType) UnmarshalJSON(b []byte) error {
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
