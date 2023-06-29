package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Stack struct {

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 资源栈的描述。可用于客户识别自己的资源栈。
	Description *string `json:"description,omitempty"`

	// 资源栈（stack）的唯一Id。  此Id由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给与的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 资源栈的状态     * `CREATION_COMPLETE` - 生成空资源栈完成，并没有任何部署     * `DEPLOYMENT_IN_PROGRESS` - 正在部署，请等待     * `DEPLOYMENT_FAILED` - 部署失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情     * `DEPLOYMENT_COMPLETE` - 部署完成     * `ROLLBACK_IN_PROGRESS` - 部署失败，正在回滚，请等待     * `ROLLBACK_FAILED` - 回滚失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情     * `ROLLBACK_COMPLETE` - 回滚完成     * `DELETION_IN_PROGRESS` - 正在删除，请等待     * `DELETION_FAILED` - 删除失败。请从status_message获取错误信息汇总，或者调用ListStackEvents获得事件详情
	Status *StackStatus `json:"status,omitempty"`

	// 资源栈的生成时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime *string `json:"create_time,omitempty"`

	// 资源栈的更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	UpdateTime *string `json:"update_time,omitempty"`

	// 在失败的时候（资源栈状态以FAILED结尾）会显示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o Stack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stack struct{}"
	}

	return strings.Join([]string{"Stack", string(data)}, " ")
}

type StackStatus struct {
	value string
}

type StackStatusEnum struct {
	CREATION_COMPLETE      StackStatus
	DEPLOYMENT_IN_PROGRESS StackStatus
	DEPLOYMENT_FAILED      StackStatus
	DEPLOYMENT_COMPLETE    StackStatus
	ROLLBACK_IN_PROGRESS   StackStatus
	ROLLBACK_FAILED        StackStatus
	ROLLBACK_COMPLETE      StackStatus
	DELETION_IN_PROGRESS   StackStatus
	DELETION_FAILED        StackStatus
}

func GetStackStatusEnum() StackStatusEnum {
	return StackStatusEnum{
		CREATION_COMPLETE: StackStatus{
			value: "CREATION_COMPLETE",
		},
		DEPLOYMENT_IN_PROGRESS: StackStatus{
			value: "DEPLOYMENT_IN_PROGRESS",
		},
		DEPLOYMENT_FAILED: StackStatus{
			value: "DEPLOYMENT_FAILED",
		},
		DEPLOYMENT_COMPLETE: StackStatus{
			value: "DEPLOYMENT_COMPLETE",
		},
		ROLLBACK_IN_PROGRESS: StackStatus{
			value: "ROLLBACK_IN_PROGRESS",
		},
		ROLLBACK_FAILED: StackStatus{
			value: "ROLLBACK_FAILED",
		},
		ROLLBACK_COMPLETE: StackStatus{
			value: "ROLLBACK_COMPLETE",
		},
		DELETION_IN_PROGRESS: StackStatus{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: StackStatus{
			value: "DELETION_FAILED",
		},
	}
}

func (c StackStatus) Value() string {
	return c.value
}

func (c StackStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackStatus) UnmarshalJSON(b []byte) error {
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
