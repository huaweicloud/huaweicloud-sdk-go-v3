package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryEventAuthorDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态。 **取值范围：** - active，激活。 - blocked，禁用。
	State *RepositoryEventAuthorDtoState `json:"state,omitempty"`

	// **参数解释：** 头像地址。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// **参数解释：** 中文名。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：** 昵称。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名。
	TenantName *string `json:"tenant_name,omitempty"`
}

func (o RepositoryEventAuthorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryEventAuthorDto struct{}"
	}

	return strings.Join([]string{"RepositoryEventAuthorDto", string(data)}, " ")
}

type RepositoryEventAuthorDtoState struct {
	value string
}

type RepositoryEventAuthorDtoStateEnum struct {
	ACTIVE  RepositoryEventAuthorDtoState
	BLOCKED RepositoryEventAuthorDtoState
}

func GetRepositoryEventAuthorDtoStateEnum() RepositoryEventAuthorDtoStateEnum {
	return RepositoryEventAuthorDtoStateEnum{
		ACTIVE: RepositoryEventAuthorDtoState{
			value: "active",
		},
		BLOCKED: RepositoryEventAuthorDtoState{
			value: "blocked",
		},
	}
}

func (c RepositoryEventAuthorDtoState) Value() string {
	return c.value
}

func (c RepositoryEventAuthorDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryEventAuthorDtoState) UnmarshalJSON(b []byte) error {
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
