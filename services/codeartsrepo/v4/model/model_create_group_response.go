package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateGroupResponse Response Object
type CreateGroupResponse struct {

	// 代码组描述信息
	Description *string `json:"description,omitempty"`

	// 代码组全名
	FullName *interface{} `json:"full_name,omitempty"`

	// 代码组id
	Id *int32 `json:"id,omitempty"`

	// 错误码
	Name *string `json:"name,omitempty"`

	// 可见性, private public
	Visibility     *CreateGroupResponseVisibility `json:"visibility,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o CreateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupResponse", string(data)}, " ")
}

type CreateGroupResponseVisibility struct {
	value string
}

type CreateGroupResponseVisibilityEnum struct {
	PUBLIC  CreateGroupResponseVisibility
	PRIVATE CreateGroupResponseVisibility
}

func GetCreateGroupResponseVisibilityEnum() CreateGroupResponseVisibilityEnum {
	return CreateGroupResponseVisibilityEnum{
		PUBLIC: CreateGroupResponseVisibility{
			value: "public",
		},
		PRIVATE: CreateGroupResponseVisibility{
			value: "private",
		},
	}
}

func (c CreateGroupResponseVisibility) Value() string {
	return c.value
}

func (c CreateGroupResponseVisibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateGroupResponseVisibility) UnmarshalJSON(b []byte) error {
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
