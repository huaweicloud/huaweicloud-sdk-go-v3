package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetUserPoliciesRequest Request Object
type SetUserPoliciesRequest struct {

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及。 **取值范围**： kafka **默认取值**： kafka
	Engine SetUserPoliciesRequestEngine `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 用户名。 **约束限制**： 不涉及。 **取值范围**： 已经创建的用户名。 **默认取值**： 不涉及。
	UserName string `json:"user_name"`

	Body *SetUserPoliciesReq `json:"body,omitempty"`
}

func (o SetUserPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserPoliciesRequest struct{}"
	}

	return strings.Join([]string{"SetUserPoliciesRequest", string(data)}, " ")
}

type SetUserPoliciesRequestEngine struct {
	value string
}

type SetUserPoliciesRequestEngineEnum struct {
	KAFKA SetUserPoliciesRequestEngine
}

func GetSetUserPoliciesRequestEngineEnum() SetUserPoliciesRequestEngineEnum {
	return SetUserPoliciesRequestEngineEnum{
		KAFKA: SetUserPoliciesRequestEngine{
			value: "kafka",
		},
	}
}

func (c SetUserPoliciesRequestEngine) Value() string {
	return c.value
}

func (c SetUserPoliciesRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetUserPoliciesRequestEngine) UnmarshalJSON(b []byte) error {
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
