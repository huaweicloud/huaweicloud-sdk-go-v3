package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackInstance struct {

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，在重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是又其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给与的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	// 资源栈集（stack_set）的名字。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	// 资源栈实例的状态  * `WAIT_IN_PROGRESS` - 资源栈实例等待操作中 * `CANCEL_COMPLETE` - 资源栈实例操作取消完成 * `OPERATION_IN_PROGRESS` - 资源栈实例操作中 * `OPERATION_FAILED` - 资源栈实例操作失败 * `INOPERABLE` - 资源栈实例不可操作 * `OPERATION_COMPLETE` - 资源栈实例操作完成
	Status *StackInstanceStatus `json:"status,omitempty"`

	// 在资源栈实例状态为`INOPERABLE`或`OPERATION_FAILED`时，会显示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`

	// 资源栈（stack）的唯一Id。  此Id由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给与的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 资源栈实例所关联的资源栈所在的租户ID
	StackDomainId *string `json:"stack_domain_id,omitempty"`

	// 最新一次部署该资源栈实例的资源栈集操作ID。  此Id由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	LatestStackSetOperationId *string `json:"latest_stack_set_operation_id,omitempty"`

	// 资源栈实例所关联的资源栈所在的区域
	Region *string `json:"region,omitempty"`

	// 资源栈实例的创建时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈实例的更新时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o StackInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackInstance struct{}"
	}

	return strings.Join([]string{"StackInstance", string(data)}, " ")
}

type StackInstanceStatus struct {
	value string
}

type StackInstanceStatusEnum struct {
	WAIT_IN_PROGRESS      StackInstanceStatus
	CANCEL_COMPLETE       StackInstanceStatus
	OPERATION_IN_PROGRESS StackInstanceStatus
	OPERATION_FAILED      StackInstanceStatus
	INOPERABLE            StackInstanceStatus
	OPERATION_COMPLETE    StackInstanceStatus
}

func GetStackInstanceStatusEnum() StackInstanceStatusEnum {
	return StackInstanceStatusEnum{
		WAIT_IN_PROGRESS: StackInstanceStatus{
			value: "WAIT_IN_PROGRESS",
		},
		CANCEL_COMPLETE: StackInstanceStatus{
			value: "CANCEL_COMPLETE",
		},
		OPERATION_IN_PROGRESS: StackInstanceStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		OPERATION_FAILED: StackInstanceStatus{
			value: "OPERATION_FAILED",
		},
		INOPERABLE: StackInstanceStatus{
			value: "INOPERABLE",
		},
		OPERATION_COMPLETE: StackInstanceStatus{
			value: "OPERATION_COMPLETE",
		},
	}
}

func (c StackInstanceStatus) Value() string {
	return c.value
}

func (c StackInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackInstanceStatus) UnmarshalJSON(b []byte) error {
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
