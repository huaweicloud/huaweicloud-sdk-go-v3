package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowTemplateFileRequest struct {
	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文

	XLanguage *ShowTemplateFileRequestXLanguage `json:"X-Language,omitempty"`
	// 模板ID，通过查询模板列表接口可获取相应的模板ID。

	TemplateId string `json:"template_id"`
	// 文件相对路径，基于当前根目录的相对文件路径，例如获取HELP.md文件内容，则文件相对路径为“template-resources/file/HELP.md”。

	FilePath string `json:"file_path"`
	// 读取文件来源，缺省值为“source-pachage”。  枚举值： - source-package: 源文件压缩包 - introduction: 说明文件

	Type *ShowTemplateFileRequestType `json:"type,omitempty"`
}

func (o ShowTemplateFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateFileRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateFileRequest", string(data)}, " ")
}

type ShowTemplateFileRequestXLanguage struct {
	value string
}

type ShowTemplateFileRequestXLanguageEnum struct {
	ZH_CN ShowTemplateFileRequestXLanguage
	EN_US ShowTemplateFileRequestXLanguage
}

func GetShowTemplateFileRequestXLanguageEnum() ShowTemplateFileRequestXLanguageEnum {
	return ShowTemplateFileRequestXLanguageEnum{
		ZH_CN: ShowTemplateFileRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowTemplateFileRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowTemplateFileRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTemplateFileRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowTemplateFileRequestType struct {
	value string
}

type ShowTemplateFileRequestTypeEnum struct {
	SOURCE_PACKAGE ShowTemplateFileRequestType
	INTRODUCTION   ShowTemplateFileRequestType
}

func GetShowTemplateFileRequestTypeEnum() ShowTemplateFileRequestTypeEnum {
	return ShowTemplateFileRequestTypeEnum{
		SOURCE_PACKAGE: ShowTemplateFileRequestType{
			value: "source-package",
		},
		INTRODUCTION: ShowTemplateFileRequestType{
			value: "introduction",
		},
	}
}

func (c ShowTemplateFileRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTemplateFileRequestType) UnmarshalJSON(b []byte) error {
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
