package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProtectedActionBasicApiDto struct {

	// **参数解释：** 事件名称。 **取值范围：push 提交,merge 合并** 字符串长度不少于1，不超过1000。
	Action *ProtectedActionBasicApiDtoAction `json:"action,omitempty"`

	// **参数解释：** 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 用户ID列表。 **约束限制：** 不涉及。 **取值范围：** Integer **默认取值：** 不涉及。
	UserIds *[]int32 `json:"user_ids,omitempty"`

	// **参数解释：** 成员组ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserTeamIds *[]int32 `json:"user_team_ids,omitempty"`

	// **参数解释：** 关联角色ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RelatedRoleIds *[]string `json:"related_role_ids,omitempty"`

	// **参数解释：** 操作选择列表。
	AdditionSwitchers *[]ForceActionEnableDto `json:"addition_switchers,omitempty"`
}

func (o ProtectedActionBasicApiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedActionBasicApiDto struct{}"
	}

	return strings.Join([]string{"ProtectedActionBasicApiDto", string(data)}, " ")
}

type ProtectedActionBasicApiDtoAction struct {
	value string
}

type ProtectedActionBasicApiDtoActionEnum struct {
	PUSH  ProtectedActionBasicApiDtoAction
	MERGE ProtectedActionBasicApiDtoAction
}

func GetProtectedActionBasicApiDtoActionEnum() ProtectedActionBasicApiDtoActionEnum {
	return ProtectedActionBasicApiDtoActionEnum{
		PUSH: ProtectedActionBasicApiDtoAction{
			value: "push",
		},
		MERGE: ProtectedActionBasicApiDtoAction{
			value: "merge",
		},
	}
}

func (c ProtectedActionBasicApiDtoAction) Value() string {
	return c.value
}

func (c ProtectedActionBasicApiDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectedActionBasicApiDtoAction) UnmarshalJSON(b []byte) error {
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
