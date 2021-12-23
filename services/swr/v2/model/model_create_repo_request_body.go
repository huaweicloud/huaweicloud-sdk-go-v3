package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateRepoRequestBody struct {
	// 镜像仓库名称。小写字母或数字开头，后面跟小写字母、数字、小数点、斜杠、下划线或中划线（其中下划线最多允许连续两个，小数点、斜杠、下划线、中划线不能直接相连），小写字母或数字结尾，1-128个字符。

	Repository string `json:"repository"`
	// 是否为公共仓库，可选值为true或false。

	IsPublic bool `json:"is_public"`
	// 仓库类型，可设置为app_server, linux, framework_app, database, lang, other, windows, arm。

	Category *CreateRepoRequestBodyCategory `json:"category,omitempty"`
	// 镜像仓库的描述信息。

	Description *string `json:"description,omitempty"`
}

func (o CreateRepoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRepoRequestBody", string(data)}, " ")
}

type CreateRepoRequestBodyCategory struct {
	value string
}

type CreateRepoRequestBodyCategoryEnum struct {
	APP_SERVER    CreateRepoRequestBodyCategory
	LINUX         CreateRepoRequestBodyCategory
	FRAMEWORK_APP CreateRepoRequestBodyCategory
	DATABASE      CreateRepoRequestBodyCategory
	LANG          CreateRepoRequestBodyCategory
	WINDOWS       CreateRepoRequestBodyCategory
	ARMS          CreateRepoRequestBodyCategory
	OTHER         CreateRepoRequestBodyCategory
}

func GetCreateRepoRequestBodyCategoryEnum() CreateRepoRequestBodyCategoryEnum {
	return CreateRepoRequestBodyCategoryEnum{
		APP_SERVER: CreateRepoRequestBodyCategory{
			value: "app_server",
		},
		LINUX: CreateRepoRequestBodyCategory{
			value: "linux",
		},
		FRAMEWORK_APP: CreateRepoRequestBodyCategory{
			value: "framework_app",
		},
		DATABASE: CreateRepoRequestBodyCategory{
			value: "database",
		},
		LANG: CreateRepoRequestBodyCategory{
			value: "lang",
		},
		WINDOWS: CreateRepoRequestBodyCategory{
			value: "windows",
		},
		ARMS: CreateRepoRequestBodyCategory{
			value: "arms",
		},
		OTHER: CreateRepoRequestBodyCategory{
			value: "other",
		},
	}
}

func (c CreateRepoRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRepoRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
