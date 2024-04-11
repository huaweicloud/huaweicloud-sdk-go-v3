package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecutionPlan struct {

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackName string `json:"stack_name"`

	// 资源栈（stack）的唯一ID。  此ID由资源编排服务在生成资源栈的时候生成，为UUID。  由于资源栈名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈，删除，再重新创建一个同名资源栈。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈就是我认为的那个，而不是其他队友删除后创建的同名资源栈。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈所对应的ID都不相同，更新不会影响ID。如果给予的stack_id和当前资源栈的ID不一致，则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 执行计划（execution_plan）的唯一Id。  此Id由资源编排服务在生成执行计划的时候生成，为UUID。  由于执行计划名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的执行计划，删除，再重新创建一个同名执行计划。  对于团队并行开发，用户可能希望确保，当前我操作的执行计划就是我认为的那个，而不是其他队友删除后创建的同名执行计划。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的执行计划所对应的ID都不相同，更新不会影响ID。如果给予的execution_plan_id和当前执行计划的ID不一致，则返回400  **注意：** * 创建执行计划后，资源编排服务持久化请求并立即返回，客户端不等待请求最终处理完成，用户无法实时感知请求处理结果 * 资源编排服务最终会将异步部署请求排队，在服务端空闲的情况下逐个处理。用户最大等待时长为1小时
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`

	// 执行计划的名称。此名字在domain_id+区域+project_id+stack_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	ExecutionPlanName string `json:"execution_plan_name"`

	// 执行计划的描述。可用于客户识别自己的执行计划。
	Description *string `json:"description,omitempty"`

	// 执行计划的状态    * `CREATION_IN_PROGRESS` - 正在创建，请等待    * `CREATION_FAILED` - 创建失败，请从status_message获取错误信息汇总    * `AVAILABLE` - 创建完成，可以调用ApplyExecutionPlan API进行执行    * `APPLY_IN_PROGRESS` - 执行中，可通过GetStackMetadata查询资源栈状态，通过ListStackEvents获取执行过程中产生的资源栈事件    * `APPLIED` - 已执行
	Status *ExecutionPlanStatus `json:"status,omitempty"`

	// 当执行计划的状态为创建失败状态（即为 `CREATION_FAILED` 时），将会展示简要的错误信息总结以供debug
	StatusMessage *string `json:"status_message,omitempty"`

	// 执行计划的生成时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime *string `json:"create_time,omitempty"`

	// 执行计划的执行时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	ApplyTime *string `json:"apply_time,omitempty"`
}

func (o ExecutionPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlan struct{}"
	}

	return strings.Join([]string{"ExecutionPlan", string(data)}, " ")
}

type ExecutionPlanStatus struct {
	value string
}

type ExecutionPlanStatusEnum struct {
	CREATION_IN_PROGRESS ExecutionPlanStatus
	CREATION_FAILED      ExecutionPlanStatus
	AVAILABLE            ExecutionPlanStatus
	APPLY_IN_PROGRESS    ExecutionPlanStatus
	APPLIED              ExecutionPlanStatus
}

func GetExecutionPlanStatusEnum() ExecutionPlanStatusEnum {
	return ExecutionPlanStatusEnum{
		CREATION_IN_PROGRESS: ExecutionPlanStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: ExecutionPlanStatus{
			value: "CREATION_FAILED",
		},
		AVAILABLE: ExecutionPlanStatus{
			value: "AVAILABLE",
		},
		APPLY_IN_PROGRESS: ExecutionPlanStatus{
			value: "APPLY_IN_PROGRESS",
		},
		APPLIED: ExecutionPlanStatus{
			value: "APPLIED",
		},
	}
}

func (c ExecutionPlanStatus) Value() string {
	return c.value
}

func (c ExecutionPlanStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionPlanStatus) UnmarshalJSON(b []byte) error {
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
