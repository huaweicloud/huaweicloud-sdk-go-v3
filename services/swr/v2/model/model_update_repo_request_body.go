package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type UpdateRepoRequestBody struct {
	// 是否为公共仓库，可选值为true或false。
	IsPublic bool `json:"is_public"`
	// 仓库类型，可设置为app_server, linux, framework_app, database, lang, other, windows, arm。
	Category *UpdateRepoRequestBodyCategory `json:"category,omitempty"`
	// 镜像仓库的描述信息。
	Description *string `json:"description,omitempty"`
}

func (o UpdateRepoRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRepoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRepoRequestBody", string(data)}, " ")
}

type UpdateRepoRequestBodyCategory struct {
	value string
}

type UpdateRepoRequestBodyCategoryEnum struct {
	APP_SERVER    UpdateRepoRequestBodyCategory
	LINUX         UpdateRepoRequestBodyCategory
	FRAMEWORK_APP UpdateRepoRequestBodyCategory
	DATABASE      UpdateRepoRequestBodyCategory
	LANG          UpdateRepoRequestBodyCategory
	WINDOWS       UpdateRepoRequestBodyCategory
	ARMS          UpdateRepoRequestBodyCategory
	OTHER         UpdateRepoRequestBodyCategory
}

func GetUpdateRepoRequestBodyCategoryEnum() UpdateRepoRequestBodyCategoryEnum {
	return UpdateRepoRequestBodyCategoryEnum{
		APP_SERVER: UpdateRepoRequestBodyCategory{
			value: "app_server",
		},
		LINUX: UpdateRepoRequestBodyCategory{
			value: "linux",
		},
		FRAMEWORK_APP: UpdateRepoRequestBodyCategory{
			value: "framework_app",
		},
		DATABASE: UpdateRepoRequestBodyCategory{
			value: "database",
		},
		LANG: UpdateRepoRequestBodyCategory{
			value: "lang",
		},
		WINDOWS: UpdateRepoRequestBodyCategory{
			value: "windows",
		},
		ARMS: UpdateRepoRequestBodyCategory{
			value: "arms",
		},
		OTHER: UpdateRepoRequestBodyCategory{
			value: "other",
		},
	}
}

func (c UpdateRepoRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateRepoRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
