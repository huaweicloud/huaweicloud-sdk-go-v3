package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateUserEmailsResponse Response Object
type UpdateUserEmailsResponse struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户长id。 **取值范围：** 字符串长度不少于1，不超过1000。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态 active 活跃用户 blocked 锁定禁用用户。
	State *UpdateUserEmailsResponseState `json:"state,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 最后活跃时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	LastActivityOn *string `json:"last_activity_on,omitempty"`

	// **参数解释：** 提交邮箱。 **取值范围：** 字符串长度不少于1，不超过1000。
	CommitEmail    *string `json:"commit_email,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserEmailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserEmailsResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserEmailsResponse", string(data)}, " ")
}

type UpdateUserEmailsResponseState struct {
	value string
}

type UpdateUserEmailsResponseStateEnum struct {
	ACTIVE  UpdateUserEmailsResponseState
	BLOCKED UpdateUserEmailsResponseState
}

func GetUpdateUserEmailsResponseStateEnum() UpdateUserEmailsResponseStateEnum {
	return UpdateUserEmailsResponseStateEnum{
		ACTIVE: UpdateUserEmailsResponseState{
			value: "active",
		},
		BLOCKED: UpdateUserEmailsResponseState{
			value: "blocked",
		},
	}
}

func (c UpdateUserEmailsResponseState) Value() string {
	return c.value
}

func (c UpdateUserEmailsResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserEmailsResponseState) UnmarshalJSON(b []byte) error {
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
