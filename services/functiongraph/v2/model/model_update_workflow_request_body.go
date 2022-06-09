package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 函数流创建body体
type UpdateWorkflowRequestBody struct {

	// 流程定义名称
	Name string `json:"name"`

	// 流程定义描述
	Description *string `json:"description,omitempty"`

	// 触发器列表
	Triggers *[]Trigger `json:"triggers,omitempty"`

	// 流程开始节点ID
	Start string `json:"start"`

	// 函数清单
	Functions []Function `json:"functions"`

	// 工作流节点清单，定义参考SleepState和OperationState
	States []OperationState `json:"states"`

	// 工作流中的常量
	Constants *interface{} `json:"constants"`

	// 重试策略清单
	Retries []Retry `json:"retries"`

	// 工作流模式，当前支持两种模式 NORMAL: 标准模式，普通模式面向普通的业务场景，支持长时间任务，支持执行历史持久化和查询，只支持异步调用 EXPRESS: 快速模式，快速模式面向业务执行时长较短，需要极致性能的场景，只支持流程执行时长低于5分钟的场景，不支持执行历史持久化，支持同步和异步调用 默认为标准模式
	Mode *UpdateWorkflowRequestBodyMode `json:"mode,omitempty"`

	ExpressConfig *ExpressConfig `json:"express_config,omitempty"`

	// 企业项目ID，在企业用户创建函数时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateWorkflowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowRequestBody", string(data)}, " ")
}

type UpdateWorkflowRequestBodyMode struct {
	value string
}

type UpdateWorkflowRequestBodyModeEnum struct {
	NORMAL  UpdateWorkflowRequestBodyMode
	EXPRESS UpdateWorkflowRequestBodyMode
}

func GetUpdateWorkflowRequestBodyModeEnum() UpdateWorkflowRequestBodyModeEnum {
	return UpdateWorkflowRequestBodyModeEnum{
		NORMAL: UpdateWorkflowRequestBodyMode{
			value: "NORMAL",
		},
		EXPRESS: UpdateWorkflowRequestBodyMode{
			value: "EXPRESS",
		},
	}
}

func (c UpdateWorkflowRequestBodyMode) Value() string {
	return c.value
}

func (c UpdateWorkflowRequestBodyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateWorkflowRequestBodyMode) UnmarshalJSON(b []byte) error {
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
