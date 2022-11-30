package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type StartPausingWorkflowExecutionsRequest struct {

	// 工作流ID，唯一标识，根据project_id和workflow_name生成。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID，xxxxxx。
	ExecutionId string `json:"execution_id"`

	// 对当前节点的操作：失败重试，失败跳过，暂停继续。 restart可重新执行失败的节点，skip可跳过失败的节点进入下个节点的执行，continue可通过暂停节点进入下一个节点。
	Action StartPausingWorkflowExecutionsRequestAction `json:"action"`

	// 当前节点的id。
	NodeId string `json:"node_id"`
}

func (o StartPausingWorkflowExecutionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPausingWorkflowExecutionsRequest struct{}"
	}

	return strings.Join([]string{"StartPausingWorkflowExecutionsRequest", string(data)}, " ")
}

type StartPausingWorkflowExecutionsRequestAction struct {
	value string
}

type StartPausingWorkflowExecutionsRequestActionEnum struct {
	RESTART  StartPausingWorkflowExecutionsRequestAction
	SKIP     StartPausingWorkflowExecutionsRequestAction
	CONTINUE StartPausingWorkflowExecutionsRequestAction
}

func GetStartPausingWorkflowExecutionsRequestActionEnum() StartPausingWorkflowExecutionsRequestActionEnum {
	return StartPausingWorkflowExecutionsRequestActionEnum{
		RESTART: StartPausingWorkflowExecutionsRequestAction{
			value: "restart",
		},
		SKIP: StartPausingWorkflowExecutionsRequestAction{
			value: "skip",
		},
		CONTINUE: StartPausingWorkflowExecutionsRequestAction{
			value: "continue",
		},
	}
}

func (c StartPausingWorkflowExecutionsRequestAction) Value() string {
	return c.value
}

func (c StartPausingWorkflowExecutionsRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartPausingWorkflowExecutionsRequestAction) UnmarshalJSON(b []byte) error {
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
