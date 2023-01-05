package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建函数流的body体
type WorkflowCreateBody struct {

	// 函数流名称
	Name string `json:"name"`

	// 函数流描述
	Description *string `json:"description,omitempty"`

	// 流程开始节点ID
	Start string `json:"start"`

	// 触发器列表
	Triggers *[]Trigger `json:"triggers,omitempty"`

	// 函数列表
	Functions []Function `json:"functions"`

	// 函数流节点清单，定义参考SleepState和OperationState
	States []OperationState `json:"states"`

	// 函数流中的常量
	Constants *interface{} `json:"constants"`

	// 重试策略清单
	Retries []Retry `json:"retries"`

	// 函数流模式，当前支持两种模式 NORMAL: 标准模式，普通模式面向普通的业务场景，支持长时间任务，支持执行历史持久化和查询，只支持异步调用 EXPRESS: 快速模式，快速模式面向业务执行时长较短，需要极致性能的场景，只支持流程执行时长低于5分钟的场景，不支持执行历史持久化，支持同步和异步调用 默认为标准模式
	Mode *WorkflowCreateBodyMode `json:"mode,omitempty"`

	ExpressConfig *ExpressConfig `json:"express_config,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否返回流数据
	EnableStreamResponse *bool `json:"enable_stream_response,omitempty"`
}

func (o WorkflowCreateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowCreateBody struct{}"
	}

	return strings.Join([]string{"WorkflowCreateBody", string(data)}, " ")
}

type WorkflowCreateBodyMode struct {
	value string
}

type WorkflowCreateBodyModeEnum struct {
	NORMAL  WorkflowCreateBodyMode
	EXPRESS WorkflowCreateBodyMode
}

func GetWorkflowCreateBodyModeEnum() WorkflowCreateBodyModeEnum {
	return WorkflowCreateBodyModeEnum{
		NORMAL: WorkflowCreateBodyMode{
			value: "NORMAL",
		},
		EXPRESS: WorkflowCreateBodyMode{
			value: "EXPRESS",
		},
	}
}

func (c WorkflowCreateBodyMode) Value() string {
	return c.value
}

func (c WorkflowCreateBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkflowCreateBodyMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
