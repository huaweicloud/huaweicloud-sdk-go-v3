package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTaskStatusReq This is a auto create Body Object
type UpdateTaskStatusReq struct {

	// 操作任务的具体动作 start:开始任务 stop:停止任务 collect_log:收集日志 test:测试 clone_test:克隆测试 restart:重新开始 sync_failed_rollback:同步失败回滚
	Operation UpdateTaskStatusReqOperation `json:"operation"`

	// 操作参数
	Param map[string]string `json:"param,omitempty"`

	// 是否切换hce
	SwitchHce *bool `json:"switch_hce,omitempty"`

	// 是否进行一致性校验
	IsNeedConsistencyCheck *bool `json:"is_need_consistency_check,omitempty"`
}

func (o UpdateTaskStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusReq struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusReq", string(data)}, " ")
}

type UpdateTaskStatusReqOperation struct {
	value string
}

type UpdateTaskStatusReqOperationEnum struct {
	START                UpdateTaskStatusReqOperation
	STOP                 UpdateTaskStatusReqOperation
	COLLECT_LOG          UpdateTaskStatusReqOperation
	TEST                 UpdateTaskStatusReqOperation
	CLONE_TEST           UpdateTaskStatusReqOperation
	RESTART              UpdateTaskStatusReqOperation
	SYNC_FAILED_ROLLBACK UpdateTaskStatusReqOperation
}

func GetUpdateTaskStatusReqOperationEnum() UpdateTaskStatusReqOperationEnum {
	return UpdateTaskStatusReqOperationEnum{
		START: UpdateTaskStatusReqOperation{
			value: "start",
		},
		STOP: UpdateTaskStatusReqOperation{
			value: "stop",
		},
		COLLECT_LOG: UpdateTaskStatusReqOperation{
			value: "collect_log",
		},
		TEST: UpdateTaskStatusReqOperation{
			value: "test",
		},
		CLONE_TEST: UpdateTaskStatusReqOperation{
			value: "clone_test",
		},
		RESTART: UpdateTaskStatusReqOperation{
			value: "restart",
		},
		SYNC_FAILED_ROLLBACK: UpdateTaskStatusReqOperation{
			value: "sync_failed_rollback",
		},
	}
}

func (c UpdateTaskStatusReqOperation) Value() string {
	return c.value
}

func (c UpdateTaskStatusReqOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTaskStatusReqOperation) UnmarshalJSON(b []byte) error {
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
