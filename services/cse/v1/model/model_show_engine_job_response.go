package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEngineJobResponse Response Object
type ShowEngineJobResponse struct {

	// 任务ID
	Id *int32 `json:"id,omitempty"`

	// 任务所属引擎ID
	EngineId *string `json:"engine_id,omitempty"`

	// 任务类型
	Type *ShowEngineJobResponseType `json:"type,omitempty"`

	// 任务描述
	Description *string `json:"description,omitempty"`

	// 任务状态
	Status *ShowEngineJobResponseStatus `json:"status,omitempty"`

	// 任务是否正在执行，0表示不在执行，1表示执行中
	Scheduling *int32 `json:"scheduling,omitempty"`

	// 任务创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 任务开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 任务执行上下文
	Context *string `json:"context,omitempty"`

	// 任务包含的处理阶段
	Tasks          *[]TaskSteps `json:"tasks,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowEngineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineJobResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineJobResponse", string(data)}, " ")
}

type ShowEngineJobResponseType struct {
	value string
}

type ShowEngineJobResponseTypeEnum struct {
	CREATE  ShowEngineJobResponseType
	DELETE  ShowEngineJobResponseType
	UPGRADE ShowEngineJobResponseType
	MODIFY  ShowEngineJobResponseType
}

func GetShowEngineJobResponseTypeEnum() ShowEngineJobResponseTypeEnum {
	return ShowEngineJobResponseTypeEnum{
		CREATE: ShowEngineJobResponseType{
			value: "Create",
		},
		DELETE: ShowEngineJobResponseType{
			value: "Delete",
		},
		UPGRADE: ShowEngineJobResponseType{
			value: "Upgrade",
		},
		MODIFY: ShowEngineJobResponseType{
			value: "Modify",
		},
	}
}

func (c ShowEngineJobResponseType) Value() string {
	return c.value
}

func (c ShowEngineJobResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineJobResponseType) UnmarshalJSON(b []byte) error {
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

type ShowEngineJobResponseStatus struct {
	value string
}

type ShowEngineJobResponseStatusEnum struct {
	INIT      ShowEngineJobResponseStatus
	EXECUTING ShowEngineJobResponseStatus
	ERROR     ShowEngineJobResponseStatus
	TIMEOUT   ShowEngineJobResponseStatus
	FINISHED  ShowEngineJobResponseStatus
}

func GetShowEngineJobResponseStatusEnum() ShowEngineJobResponseStatusEnum {
	return ShowEngineJobResponseStatusEnum{
		INIT: ShowEngineJobResponseStatus{
			value: "Init",
		},
		EXECUTING: ShowEngineJobResponseStatus{
			value: "Executing",
		},
		ERROR: ShowEngineJobResponseStatus{
			value: "Error",
		},
		TIMEOUT: ShowEngineJobResponseStatus{
			value: "Timeout",
		},
		FINISHED: ShowEngineJobResponseStatus{
			value: "Finished",
		},
	}
}

func (c ShowEngineJobResponseStatus) Value() string {
	return c.value
}

func (c ShowEngineJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineJobResponseStatus) UnmarshalJSON(b []byte) error {
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
