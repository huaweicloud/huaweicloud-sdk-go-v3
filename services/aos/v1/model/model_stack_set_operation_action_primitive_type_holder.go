package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackSetOperationActionPrimitiveTypeHolder struct {

	// 用户当前的操作   * `CREATE_STACK_INSTANCES` - 创建资源栈实例   * `DELETE_STACK_INSTANCES` - 删除资源栈实例   * `DEPLOY_STACK_SET` - 部署资源栈集   * `DEPLOY_STACK_INSTANCES` - 部署资源栈实例   * `UPDATE_STACK_INSTANCES` - 更新资源栈实例
	Action *StackSetOperationActionPrimitiveTypeHolderAction `json:"action,omitempty"`
}

func (o StackSetOperationActionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperationActionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetOperationActionPrimitiveTypeHolder", string(data)}, " ")
}

type StackSetOperationActionPrimitiveTypeHolderAction struct {
	value string
}

type StackSetOperationActionPrimitiveTypeHolderActionEnum struct {
	CREATE_STACK_INSTANCES StackSetOperationActionPrimitiveTypeHolderAction
	DELETE_STACK_INSTANCES StackSetOperationActionPrimitiveTypeHolderAction
	DEPLOY_STACK_SET       StackSetOperationActionPrimitiveTypeHolderAction
	DEPLOY_STACK_INSTANCES StackSetOperationActionPrimitiveTypeHolderAction
	UPDATE_STACK_INSTANCES StackSetOperationActionPrimitiveTypeHolderAction
}

func GetStackSetOperationActionPrimitiveTypeHolderActionEnum() StackSetOperationActionPrimitiveTypeHolderActionEnum {
	return StackSetOperationActionPrimitiveTypeHolderActionEnum{
		CREATE_STACK_INSTANCES: StackSetOperationActionPrimitiveTypeHolderAction{
			value: "CREATE_STACK_INSTANCES",
		},
		DELETE_STACK_INSTANCES: StackSetOperationActionPrimitiveTypeHolderAction{
			value: "DELETE_STACK_INSTANCES",
		},
		DEPLOY_STACK_SET: StackSetOperationActionPrimitiveTypeHolderAction{
			value: "DEPLOY_STACK_SET",
		},
		DEPLOY_STACK_INSTANCES: StackSetOperationActionPrimitiveTypeHolderAction{
			value: "DEPLOY_STACK_INSTANCES",
		},
		UPDATE_STACK_INSTANCES: StackSetOperationActionPrimitiveTypeHolderAction{
			value: "UPDATE_STACK_INSTANCES",
		},
	}
}

func (c StackSetOperationActionPrimitiveTypeHolderAction) Value() string {
	return c.value
}

func (c StackSetOperationActionPrimitiveTypeHolderAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetOperationActionPrimitiveTypeHolderAction) UnmarshalJSON(b []byte) error {
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
