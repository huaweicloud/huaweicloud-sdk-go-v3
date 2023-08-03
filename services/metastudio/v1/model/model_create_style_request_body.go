package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateStyleRequestBody 创建风格请求
type CreateStyleRequestBody struct {

	// 数字人风格化名称
	Name string `json:"name"`

	// 数字人风格化描述
	Description *string `json:"description,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 性别
	Sex *CreateStyleRequestBodySex `json:"sex,omitempty"`

	// 标签。单个标签16字节，多个用逗号分隔，最多50个。
	Tags *[]string `json:"tags,omitempty"`

	// 风格化素材资产组合。
	StyleAssets *[]StyleAssetItem `json:"style_assets,omitempty"`

	ExtraMeta *StyleExtraMeta `json:"extra_meta,omitempty"`
}

func (o CreateStyleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStyleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateStyleRequestBody", string(data)}, " ")
}

type CreateStyleRequestBodySex struct {
	value string
}

type CreateStyleRequestBodySexEnum struct {
	UNKNOW CreateStyleRequestBodySex
	MALE   CreateStyleRequestBodySex
	FEMALE CreateStyleRequestBodySex
}

func GetCreateStyleRequestBodySexEnum() CreateStyleRequestBodySexEnum {
	return CreateStyleRequestBodySexEnum{
		UNKNOW: CreateStyleRequestBodySex{
			value: "UNKNOW",
		},
		MALE: CreateStyleRequestBodySex{
			value: "MALE",
		},
		FEMALE: CreateStyleRequestBodySex{
			value: "FEMALE",
		},
	}
}

func (c CreateStyleRequestBodySex) Value() string {
	return c.value
}

func (c CreateStyleRequestBodySex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateStyleRequestBodySex) UnmarshalJSON(b []byte) error {
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
