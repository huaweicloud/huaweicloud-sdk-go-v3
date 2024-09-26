package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateGroupRequestBody 创建代码组请求体
type CreateGroupRequestBody struct {

	// 代码组名称
	Name *string `json:"name,omitempty"`

	// 父级代码组id, 不传默认在项目下创建代码组
	ParentId *int32 `json:"parent_id,omitempty"`

	// 可见性, private public
	Visibility *CreateGroupRequestBodyVisibility `json:"visibility,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o CreateGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGroupRequestBody", string(data)}, " ")
}

type CreateGroupRequestBodyVisibility struct {
	value string
}

type CreateGroupRequestBodyVisibilityEnum struct {
	PUBLIC  CreateGroupRequestBodyVisibility
	PRIVATE CreateGroupRequestBodyVisibility
}

func GetCreateGroupRequestBodyVisibilityEnum() CreateGroupRequestBodyVisibilityEnum {
	return CreateGroupRequestBodyVisibilityEnum{
		PUBLIC: CreateGroupRequestBodyVisibility{
			value: "public",
		},
		PRIVATE: CreateGroupRequestBodyVisibility{
			value: "private",
		},
	}
}

func (c CreateGroupRequestBodyVisibility) Value() string {
	return c.value
}

func (c CreateGroupRequestBodyVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGroupRequestBodyVisibility) UnmarshalJSON(b []byte) error {
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
