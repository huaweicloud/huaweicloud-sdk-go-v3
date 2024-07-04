package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTaskStatusReq This is a auto create Body Object
type UpdateTaskStatusReq struct {

	// 操作任务的具体动作 start:开始任务 stop:停止任务 test:测试 clone_test:克隆测试 restart:重新开始 network_check:网络质量检测
	Operation UpdateTaskStatusReqOperation `json:"operation"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

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
	START         UpdateTaskStatusReqOperation
	STOP          UpdateTaskStatusReqOperation
	TEST          UpdateTaskStatusReqOperation
	CLONE_TEST    UpdateTaskStatusReqOperation
	RESTART       UpdateTaskStatusReqOperation
	NETWORK_CHECK UpdateTaskStatusReqOperation
}

func GetUpdateTaskStatusReqOperationEnum() UpdateTaskStatusReqOperationEnum {
	return UpdateTaskStatusReqOperationEnum{
		START: UpdateTaskStatusReqOperation{
			value: "start",
		},
		STOP: UpdateTaskStatusReqOperation{
			value: "stop",
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
		NETWORK_CHECK: UpdateTaskStatusReqOperation{
			value: "network_check",
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
