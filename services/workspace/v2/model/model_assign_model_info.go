package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssignModelInfo 批量分配桌面策略。
type AssignModelInfo struct {

	// 分配的维度，当前支持  - USER: “用户为维度” - DESKTOP:“桌面为维度”两种。
	AssignDimension *AssignModelInfoAssignDimension `json:"assign_dimension,omitempty"`

	// 优先分配的策略，策略名为规则为{维度}_{子策略}  - USER_NO_DESKTOP:  用户维度-无桌面 - USER_FIXED_DESKTOP_NUM: 用户维度-桌面个数固定 - DESKTOP_ASSIGN_USER_PRIORITY： 桌面维度-用户优先 - DESKTOP_ASSIGN_FIXED_USER： 桌面维度-固定用户 - DESKTOP_ASSIGN_USERS_OR_GROUPS： 桌面维度-为每台桌面分配所选的所有用户（组） - FIXED_RELATION: 使用参数中的固定分配关系
	PriorityStrategy *AssignModelInfoPriorityStrategy `json:"priority_strategy,omitempty"`

	// 每个桌面自动分配的用户数，当子策略为 DESKTOP_ASSIGN_FIXED_USER 必填。
	DesktopAssiginUserNum *int32 `json:"desktop_assigin_user_num,omitempty"`

	// 每个用户自动分配桌面数，当子策略为 USER_NO_DESKTOP、USER_FIXED_DESKTOP_NUM必填。
	UserAssiginDesktopNum *int32 `json:"user_assigin_desktop_num,omitempty"`

	// 策略id，用于指定生成桌面名称策略，如果指定了桌面名称则优先使用指定的桌面名称。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`
}

func (o AssignModelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignModelInfo struct{}"
	}

	return strings.Join([]string{"AssignModelInfo", string(data)}, " ")
}

type AssignModelInfoAssignDimension struct {
	value string
}

type AssignModelInfoAssignDimensionEnum struct {
	USER    AssignModelInfoAssignDimension
	DESKTOP AssignModelInfoAssignDimension
}

func GetAssignModelInfoAssignDimensionEnum() AssignModelInfoAssignDimensionEnum {
	return AssignModelInfoAssignDimensionEnum{
		USER: AssignModelInfoAssignDimension{
			value: "USER",
		},
		DESKTOP: AssignModelInfoAssignDimension{
			value: "DESKTOP",
		},
	}
}

func (c AssignModelInfoAssignDimension) Value() string {
	return c.value
}

func (c AssignModelInfoAssignDimension) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssignModelInfoAssignDimension) UnmarshalJSON(b []byte) error {
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

type AssignModelInfoPriorityStrategy struct {
	value string
}

type AssignModelInfoPriorityStrategyEnum struct {
	USER_NO_DESKTOP                AssignModelInfoPriorityStrategy
	USER_FIXED_DESKTOP_NUM         AssignModelInfoPriorityStrategy
	DESKTOP_ASSIGN_USER_PRIORITY   AssignModelInfoPriorityStrategy
	DESKTOP_ASSIGN_FIXED_USER      AssignModelInfoPriorityStrategy
	DESKTOP_ASSIGN_USERS_OR_GROUPS AssignModelInfoPriorityStrategy
	FIXED_RELATION                 AssignModelInfoPriorityStrategy
}

func GetAssignModelInfoPriorityStrategyEnum() AssignModelInfoPriorityStrategyEnum {
	return AssignModelInfoPriorityStrategyEnum{
		USER_NO_DESKTOP: AssignModelInfoPriorityStrategy{
			value: "USER_NO_DESKTOP",
		},
		USER_FIXED_DESKTOP_NUM: AssignModelInfoPriorityStrategy{
			value: "USER_FIXED_DESKTOP_NUM",
		},
		DESKTOP_ASSIGN_USER_PRIORITY: AssignModelInfoPriorityStrategy{
			value: "DESKTOP_ASSIGN_USER_PRIORITY",
		},
		DESKTOP_ASSIGN_FIXED_USER: AssignModelInfoPriorityStrategy{
			value: "DESKTOP_ASSIGN_FIXED_USER",
		},
		DESKTOP_ASSIGN_USERS_OR_GROUPS: AssignModelInfoPriorityStrategy{
			value: "DESKTOP_ASSIGN_USERS_OR_GROUPS",
		},
		FIXED_RELATION: AssignModelInfoPriorityStrategy{
			value: "FIXED_RELATION",
		},
	}
}

func (c AssignModelInfoPriorityStrategy) Value() string {
	return c.value
}

func (c AssignModelInfoPriorityStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssignModelInfoPriorityStrategy) UnmarshalJSON(b []byte) error {
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
