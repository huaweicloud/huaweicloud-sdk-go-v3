package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUserEmailsResponse Response Object
type ShowUserEmailsResponse struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户长id。 **取值范围：** 字符串长度不少于1，不超过1000。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态 active 活跃用户 blocked 锁定禁用用户。
	State *ShowUserEmailsResponseState `json:"state,omitempty"`

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

func (o ShowUserEmailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserEmailsResponse struct{}"
	}

	return strings.Join([]string{"ShowUserEmailsResponse", string(data)}, " ")
}

type ShowUserEmailsResponseState struct {
	value string
}

type ShowUserEmailsResponseStateEnum struct {
	ACTIVE  ShowUserEmailsResponseState
	BLOCKED ShowUserEmailsResponseState
}

func GetShowUserEmailsResponseStateEnum() ShowUserEmailsResponseStateEnum {
	return ShowUserEmailsResponseStateEnum{
		ACTIVE: ShowUserEmailsResponseState{
			value: "active",
		},
		BLOCKED: ShowUserEmailsResponseState{
			value: "blocked",
		},
	}
}

func (c ShowUserEmailsResponseState) Value() string {
	return c.value
}

func (c ShowUserEmailsResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserEmailsResponseState) UnmarshalJSON(b []byte) error {
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
