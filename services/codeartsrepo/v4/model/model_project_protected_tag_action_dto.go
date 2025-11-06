package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProjectProtectedTagActionDto struct {

	// **参数解释：** 事件名称。 **取值范围：read 查询,create-delete 增删,create 创建** 字符串长度不少于1，不超过1000。
	Action *ProjectProtectedTagActionDtoAction `json:"action,omitempty"`

	// **参数解释：** 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 用户ID列表。 **约束限制：** 不涉及。 **取值范围：** Integer **默认取值：** 不涉及。
	UserIds *[]int32 `json:"user_ids,omitempty"`

	// **参数解释：** 用户name列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserNames *[]string `json:"user_names,omitempty"`

	// **参数解释：** 成员组ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserTeamIds *[]int32 `json:"user_team_ids,omitempty"`

	// **参数解释：** 成员组name列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserTeamNames *[]string `json:"user_team_names,omitempty"`

	// **参数解释：** 关联角色ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RelatedRoleIds *[]string `json:"related_role_ids,omitempty"`
}

func (o ProjectProtectedTagActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectProtectedTagActionDto struct{}"
	}

	return strings.Join([]string{"ProjectProtectedTagActionDto", string(data)}, " ")
}

type ProjectProtectedTagActionDtoAction struct {
	value string
}

type ProjectProtectedTagActionDtoActionEnum struct {
	READ          ProjectProtectedTagActionDtoAction
	CREATE_DELETE ProjectProtectedTagActionDtoAction
	CREATE        ProjectProtectedTagActionDtoAction
}

func GetProjectProtectedTagActionDtoActionEnum() ProjectProtectedTagActionDtoActionEnum {
	return ProjectProtectedTagActionDtoActionEnum{
		READ: ProjectProtectedTagActionDtoAction{
			value: "read",
		},
		CREATE_DELETE: ProjectProtectedTagActionDtoAction{
			value: "create-delete",
		},
		CREATE: ProjectProtectedTagActionDtoAction{
			value: "create",
		},
	}
}

func (c ProjectProtectedTagActionDtoAction) Value() string {
	return c.value
}

func (c ProjectProtectedTagActionDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectProtectedTagActionDtoAction) UnmarshalJSON(b []byte) error {
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
