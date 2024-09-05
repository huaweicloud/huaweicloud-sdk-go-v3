package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPolicyAssignmentResponse Response Object
type ShowPolicyAssignmentResponse struct {

	// 规则类型，包括预定义合规规则(builtin)和用户自定义合规规则(custom)
	PolicyAssignmentType *ShowPolicyAssignmentResponsePolicyAssignmentType `json:"policy_assignment_type,omitempty"`

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 规则名字
	Name *string `json:"name,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty"`

	// 触发周期值，可选值：One_Hour, Three_Hours, Six_Hours, Twelve_Hours, TwentyFour_Hours
	Period *string `json:"period,omitempty"`

	// 规则状态
	State *string `json:"state,omitempty"`

	// 规则创建时间
	Created *string `json:"created,omitempty"`

	// 规则更新时间
	Updated *string `json:"updated,omitempty"`

	// 规则的策略ID
	PolicyDefinitionId *string `json:"policy_definition_id,omitempty"`

	CustomPolicy *CustomPolicy `json:"custom_policy,omitempty"`

	// 规则参数
	Parameters map[string]PolicyParameterValue `json:"parameters,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 规则的创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 合规规则修正方式。
	TargetType *string `json:"target_type,omitempty"`

	// 修正执行的目标id。
	TargetId       *string `json:"target_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyAssignmentResponse", string(data)}, " ")
}

type ShowPolicyAssignmentResponsePolicyAssignmentType struct {
	value string
}

type ShowPolicyAssignmentResponsePolicyAssignmentTypeEnum struct {
	BUILTIN ShowPolicyAssignmentResponsePolicyAssignmentType
	CUSTOM  ShowPolicyAssignmentResponsePolicyAssignmentType
}

func GetShowPolicyAssignmentResponsePolicyAssignmentTypeEnum() ShowPolicyAssignmentResponsePolicyAssignmentTypeEnum {
	return ShowPolicyAssignmentResponsePolicyAssignmentTypeEnum{
		BUILTIN: ShowPolicyAssignmentResponsePolicyAssignmentType{
			value: "builtin",
		},
		CUSTOM: ShowPolicyAssignmentResponsePolicyAssignmentType{
			value: "custom",
		},
	}
}

func (c ShowPolicyAssignmentResponsePolicyAssignmentType) Value() string {
	return c.value
}

func (c ShowPolicyAssignmentResponsePolicyAssignmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPolicyAssignmentResponsePolicyAssignmentType) UnmarshalJSON(b []byte) error {
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
