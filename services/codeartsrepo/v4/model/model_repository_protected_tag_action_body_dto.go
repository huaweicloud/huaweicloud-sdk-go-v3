package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepositoryProtectedTagActionBodyDto struct {

	// **参数解释：** 是否启用。 **约束限制：** 不涉及。 **取值范围：** - true，开启规则限制 - false，关闭规则限制 **默认取值：** 不涉及。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 用户ID列表。 **约束限制：** 不涉及。 **取值范围：** Integer **默认取值：** 不涉及。
	UserIds *[]int32 `json:"user_ids,omitempty"`

	// **参数解释：** 成员组ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserTeamIds *[]int32 `json:"user_team_ids,omitempty"`

	// **参数解释：** 关联角色ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RelatedRoleIds *[]string `json:"related_role_ids,omitempty"`

	// **参数解释：** 事件名称。 **约束限制：** 不涉及 **取值范围：** - create，创建。 **默认取值：** 不涉及。
	Action *RepositoryProtectedTagActionBodyDtoAction `json:"action,omitempty"`
}

func (o RepositoryProtectedTagActionBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryProtectedTagActionBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryProtectedTagActionBodyDto", string(data)}, " ")
}

type RepositoryProtectedTagActionBodyDtoAction struct {
	value string
}

type RepositoryProtectedTagActionBodyDtoActionEnum struct {
	CREATE RepositoryProtectedTagActionBodyDtoAction
}

func GetRepositoryProtectedTagActionBodyDtoActionEnum() RepositoryProtectedTagActionBodyDtoActionEnum {
	return RepositoryProtectedTagActionBodyDtoActionEnum{
		CREATE: RepositoryProtectedTagActionBodyDtoAction{
			value: "create",
		},
	}
}

func (c RepositoryProtectedTagActionBodyDtoAction) Value() string {
	return c.value
}

func (c RepositoryProtectedTagActionBodyDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryProtectedTagActionBodyDtoAction) UnmarshalJSON(b []byte) error {
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
