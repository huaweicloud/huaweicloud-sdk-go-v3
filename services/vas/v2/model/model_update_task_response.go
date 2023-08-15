package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateTaskResponse Response Object
type UpdateTaskResponse struct {

	// 作业ID
	Id string `json:"id"`

	// 作业的名称
	Name string `json:"name"`

	// 作业创建者的用户名
	Creator string `json:"creator"`

	// 作业创建者的项目ID
	ProjectId string `json:"project_id"`

	// 作业的描述
	Description *string `json:"description,omitempty"`

	// 作业对应服务的名称
	ServiceName string `json:"service_name"`

	// 作业对应服务的版本号
	ServiceVersion string `json:"service_version"`

	ServiceTitle *TaskDetailsServiceTitle `json:"service_title"`

	// 仅边缘作业会出现，作业运行所在的边缘运行池ID
	EdgePoolId *string `json:"edge_pool_id,omitempty"`

	// 作业指定的算法能力包包周期订单ID
	ResourceOrderId *string `json:"resource_order_id,omitempty"`

	// 作业创建的时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 作业最近一次状态更新的时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 作业当前的状态，分别为PENDING（等待中），RECOVERING（恢复中），STARTING（启动中），UPGRADING（升级中），CREATE_FAILED（创建失败），START_FAILED（启动失败），RUNNING（运行中），STOPPING（停止中），STOPPED（已停止），ABNORMAL（异常），SUCCEEDED（运行成功），FAILED（运行失败），DELETING（删除中），FREEZING（冻结中），FROZEN（已冻结）
	State UpdateTaskResponseState `json:"state"`

	// 作业状态的详情信息，仅部分状态会有详情信息
	Status *string `json:"status,omitempty"`

	Error *TaskDetailsError `json:"error,omitempty"`

	// 计划任务的状态，分别为ACTIVATED（激活），INACTIVATED（未激活）
	TimingStatus *UpdateTaskResponseTimingStatus `json:"timing_status,omitempty"`

	Timing *TaskTiming `json:"timing,omitempty"`

	Input *TaskInput `json:"input"`

	Output *TaskOutputForDisplay `json:"output"`

	ServiceConfig  *TaskServiceConfig `json:"service_config,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskResponse", string(data)}, " ")
}

type UpdateTaskResponseState struct {
	value string
}

type UpdateTaskResponseStateEnum struct {
	PENDING       UpdateTaskResponseState
	RECOVERING    UpdateTaskResponseState
	STARTING      UpdateTaskResponseState
	UPGRADING     UpdateTaskResponseState
	CREATE_FAILED UpdateTaskResponseState
	START_FAILED  UpdateTaskResponseState
	RUNNING       UpdateTaskResponseState
	STOPPING      UpdateTaskResponseState
	STOPPED       UpdateTaskResponseState
	ABNORMAL      UpdateTaskResponseState
	SUCCEEDED     UpdateTaskResponseState
	FAILED        UpdateTaskResponseState
	DELETING      UpdateTaskResponseState
	FREEZING      UpdateTaskResponseState
	FROZEN        UpdateTaskResponseState
}

func GetUpdateTaskResponseStateEnum() UpdateTaskResponseStateEnum {
	return UpdateTaskResponseStateEnum{
		PENDING: UpdateTaskResponseState{
			value: "PENDING",
		},
		RECOVERING: UpdateTaskResponseState{
			value: "RECOVERING",
		},
		STARTING: UpdateTaskResponseState{
			value: "STARTING",
		},
		UPGRADING: UpdateTaskResponseState{
			value: "UPGRADING",
		},
		CREATE_FAILED: UpdateTaskResponseState{
			value: "CREATE_FAILED",
		},
		START_FAILED: UpdateTaskResponseState{
			value: "START_FAILED",
		},
		RUNNING: UpdateTaskResponseState{
			value: "RUNNING",
		},
		STOPPING: UpdateTaskResponseState{
			value: "STOPPING",
		},
		STOPPED: UpdateTaskResponseState{
			value: "STOPPED",
		},
		ABNORMAL: UpdateTaskResponseState{
			value: "ABNORMAL",
		},
		SUCCEEDED: UpdateTaskResponseState{
			value: "SUCCEEDED",
		},
		FAILED: UpdateTaskResponseState{
			value: "FAILED",
		},
		DELETING: UpdateTaskResponseState{
			value: "DELETING",
		},
		FREEZING: UpdateTaskResponseState{
			value: "FREEZING",
		},
		FROZEN: UpdateTaskResponseState{
			value: "FROZEN",
		},
	}
}

func (c UpdateTaskResponseState) Value() string {
	return c.value
}

func (c UpdateTaskResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTaskResponseState) UnmarshalJSON(b []byte) error {
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

type UpdateTaskResponseTimingStatus struct {
	value string
}

type UpdateTaskResponseTimingStatusEnum struct {
	ACTIVATED   UpdateTaskResponseTimingStatus
	INACTIVATED UpdateTaskResponseTimingStatus
}

func GetUpdateTaskResponseTimingStatusEnum() UpdateTaskResponseTimingStatusEnum {
	return UpdateTaskResponseTimingStatusEnum{
		ACTIVATED: UpdateTaskResponseTimingStatus{
			value: "ACTIVATED",
		},
		INACTIVATED: UpdateTaskResponseTimingStatus{
			value: "INACTIVATED",
		},
	}
}

func (c UpdateTaskResponseTimingStatus) Value() string {
	return c.value
}

func (c UpdateTaskResponseTimingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTaskResponseTimingStatus) UnmarshalJSON(b []byte) error {
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
