package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackSetOperation struct {

	// 资源栈集操作Id。  此ID由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	OperationId *string `json:"operation_id,omitempty"`

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给予的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 资源栈集（stack_set）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	// 用户当前的操作   * `CREATE_STACK_INSTANCES` - 创建资源栈实例   * `DELETE_STACK_INSTANCES` - 删除资源栈实例   * `DEPLOY_STACK_SET` - 部署资源栈集   * `DEPLOY_STACK_INSTANCES` - 部署资源栈实例   * `UPDATE_STACK_INSTANCES` - 更新资源栈实例
	Action *StackSetOperationAction `json:"action,omitempty"`

	// 资源栈集操作状态   * `QUEUE_IN_PROGRESS` - 正在排队   * `OPERATION_IN_PROGRESS` - 正在操作   * `OPERATION_COMPLETE` - 操作完成   * `OPERATION_FAILED` - 操作失败   * `STOP_IN_PROGRESS` - 正在停止   * `STOP_COMPLETE` - 停止完成   * `STOP_FAILED` - 停止失败
	Status *StackSetOperationStatus `json:"status,omitempty"`

	// 资源栈集操作失败时会展示此次操作失败的原因，例如，资源栈实例部署或删除失败个数超过上限或资源栈集操作超时。  如果需要查看详细失败信息，可通过ListStackInstances API获取查看资源栈实例的status_message。
	StatusMessage *string `json:"status_message,omitempty"`

	// 资源栈集操作的创建时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈集操作的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o StackSetOperation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperation struct{}"
	}

	return strings.Join([]string{"StackSetOperation", string(data)}, " ")
}

type StackSetOperationAction struct {
	value string
}

type StackSetOperationActionEnum struct {
	CREATE_STACK_INSTANCES StackSetOperationAction
	DELETE_STACK_INSTANCES StackSetOperationAction
	DEPLOY_STACK_SET       StackSetOperationAction
	DEPLOY_STACK_INSTANCES StackSetOperationAction
	UPDATE_STACK_INSTANCES StackSetOperationAction
}

func GetStackSetOperationActionEnum() StackSetOperationActionEnum {
	return StackSetOperationActionEnum{
		CREATE_STACK_INSTANCES: StackSetOperationAction{
			value: "CREATE_STACK_INSTANCES",
		},
		DELETE_STACK_INSTANCES: StackSetOperationAction{
			value: "DELETE_STACK_INSTANCES",
		},
		DEPLOY_STACK_SET: StackSetOperationAction{
			value: "DEPLOY_STACK_SET",
		},
		DEPLOY_STACK_INSTANCES: StackSetOperationAction{
			value: "DEPLOY_STACK_INSTANCES",
		},
		UPDATE_STACK_INSTANCES: StackSetOperationAction{
			value: "UPDATE_STACK_INSTANCES",
		},
	}
}

func (c StackSetOperationAction) Value() string {
	return c.value
}

func (c StackSetOperationAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetOperationAction) UnmarshalJSON(b []byte) error {
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

type StackSetOperationStatus struct {
	value string
}

type StackSetOperationStatusEnum struct {
	QUEUE_IN_PROGRESS     StackSetOperationStatus
	OPERATION_IN_PROGRESS StackSetOperationStatus
	OPERATION_COMPLETE    StackSetOperationStatus
	OPERATION_FAILED      StackSetOperationStatus
	STOP_IN_PROGRESS      StackSetOperationStatus
	STOP_COMPLETE         StackSetOperationStatus
	STOP_FAILED           StackSetOperationStatus
}

func GetStackSetOperationStatusEnum() StackSetOperationStatusEnum {
	return StackSetOperationStatusEnum{
		QUEUE_IN_PROGRESS: StackSetOperationStatus{
			value: "QUEUE_IN_PROGRESS",
		},
		OPERATION_IN_PROGRESS: StackSetOperationStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		OPERATION_COMPLETE: StackSetOperationStatus{
			value: "OPERATION_COMPLETE",
		},
		OPERATION_FAILED: StackSetOperationStatus{
			value: "OPERATION_FAILED",
		},
		STOP_IN_PROGRESS: StackSetOperationStatus{
			value: "STOP_IN_PROGRESS",
		},
		STOP_COMPLETE: StackSetOperationStatus{
			value: "STOP_COMPLETE",
		},
		STOP_FAILED: StackSetOperationStatus{
			value: "STOP_FAILED",
		},
	}
}

func (c StackSetOperationStatus) Value() string {
	return c.value
}

func (c StackSetOperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetOperationStatus) UnmarshalJSON(b []byte) error {
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
