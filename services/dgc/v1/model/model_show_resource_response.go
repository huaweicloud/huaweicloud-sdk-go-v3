package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowResourceResponse Response Object
type ShowResourceResponse struct {

	// 资源id
	Id *string `json:"id,omitempty"`

	// 资源名称，只能包含英文字母、数字、中文字符、下划线或中划线。
	Name *string `json:"name,omitempty"`

	// 资源类型:   - archive: 压缩包   - file: 文件   - jar: jar文件   - pyFile：python文件
	Type *ShowResourceResponseType `json:"type,omitempty"`

	// 资源文件所在OBS路径
	Location *string `json:"location,omitempty"`

	// 主Jar包所依赖的JAR包、properties文件
	DependFiles *[]string `json:"dependFiles,omitempty"`

	// 资源描述
	Desc *string `json:"desc,omitempty"`

	// 资源所在目录
	Directory *string `json:"directory,omitempty"`

	// 主Jar包所依赖的JAR包、properties文件。同时存在dependFiles和dependPackages时，优先解析该字段。
	DependPackages *[]DependPackage `json:"dependPackages,omitempty"`

	// 通过jar包名称查询相关的job
	JobRelation    *[]Relation `json:"jobRelation,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowResourceResponse) String() string {
	data, err := utils.Marshal(o)
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
	PY_FILE ShowResourceResponseType
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
		PY_FILE: ShowResourceResponseType{
			value: "pyFile",
		},
	}
}

func (c ShowResourceResponseType) Value() string {
	return c.value
}

func (c ShowResourceResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceResponseType) UnmarshalJSON(b []byte) error {
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
