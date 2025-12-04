package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SetUserPoliciesReq struct {

	// **参数解释**： 操作名称。 **约束限制**： 不涉及。 **取值范围**： - create：提交创建用户策略任务。 - delete：提交删除用户策略任务。 **默认取值**： 不涉及。
	Action *SetUserPoliciesReqAction `json:"action,omitempty"`

	// **参数解释**： 用户权限列表。 **约束限制**： 不涉及。
	Policies *[]UserPolicyEntity `json:"policies,omitempty"`
}

func (o SetUserPoliciesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserPoliciesReq struct{}"
	}

	return strings.Join([]string{"SetUserPoliciesReq", string(data)}, " ")
}

type SetUserPoliciesReqAction struct {
	value string
}

type SetUserPoliciesReqActionEnum struct {
	CREATE SetUserPoliciesReqAction
	DELETE SetUserPoliciesReqAction
}

func GetSetUserPoliciesReqActionEnum() SetUserPoliciesReqActionEnum {
	return SetUserPoliciesReqActionEnum{
		CREATE: SetUserPoliciesReqAction{
			value: "create",
		},
		DELETE: SetUserPoliciesReqAction{
			value: "delete",
		},
	}
}

func (c SetUserPoliciesReqAction) Value() string {
	return c.value
}

func (c SetUserPoliciesReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetUserPoliciesReqAction) UnmarshalJSON(b []byte) error {
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
