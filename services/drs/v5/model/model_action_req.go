package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ActionReq 测试连接、预检查、启动、暂停、续传、重置、对比、结束等操作任务请求体。
type ActionReq struct {

	// 任务ID (对比任务相关操作，多任务场景传父任务详情返回的master_job_id)，批量操作时必填
	JobId *string `json:"job_id,omitempty"`

	// 操作任务动作名称。取值： - network：测试连接源库/目标库。 - precheck：执行预检查。 - start：启动任务。 - stop：暂停任务。 - restart：重试任务。 - reset：重置任务。 - terminate：结束任务。 - skip_precheck：跳过预检查。 - create_compare：创建对比任务。 - cancel_compare：取消对比任务。
	ActionName ActionReqActionName `json:"action_name"`

	ActionParams *ActionParams `json:"action_params,omitempty"`
}

func (o ActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionReq struct{}"
	}

	return strings.Join([]string{"ActionReq", string(data)}, " ")
}

type ActionReqActionName struct {
	value string
}

type ActionReqActionNameEnum struct {
	NETWORK        ActionReqActionName
	PRECHECK       ActionReqActionName
	START          ActionReqActionName
	STOP           ActionReqActionName
	RESTART        ActionReqActionName
	RESET          ActionReqActionName
	TERMINATE      ActionReqActionName
	SKIP_PRECHECK  ActionReqActionName
	CREATE_COMPARE ActionReqActionName
	CANCEL_COMPARE ActionReqActionName
}

func GetActionReqActionNameEnum() ActionReqActionNameEnum {
	return ActionReqActionNameEnum{
		NETWORK: ActionReqActionName{
			value: "network",
		},
		PRECHECK: ActionReqActionName{
			value: "precheck",
		},
		START: ActionReqActionName{
			value: "start",
		},
		STOP: ActionReqActionName{
			value: "stop",
		},
		RESTART: ActionReqActionName{
			value: "restart",
		},
		RESET: ActionReqActionName{
			value: "reset",
		},
		TERMINATE: ActionReqActionName{
			value: "terminate",
		},
		SKIP_PRECHECK: ActionReqActionName{
			value: "skip_precheck",
		},
		CREATE_COMPARE: ActionReqActionName{
			value: "create_compare",
		},
		CANCEL_COMPARE: ActionReqActionName{
			value: "cancel_compare",
		},
	}
}

func (c ActionReqActionName) Value() string {
	return c.value
}

func (c ActionReqActionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ActionReqActionName) UnmarshalJSON(b []byte) error {
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
