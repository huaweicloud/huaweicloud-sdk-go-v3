package model

import (
	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowTaskResponse struct {
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

	State ShowTaskResponseState `json:"state"`
	// 作业状态的详情信息，仅部分状态会有详情信息

	Status *string `json:"status,omitempty"`

	Error *TaskDetailsError `json:"error,omitempty"`
	// 计划任务的状态，分别为ACTIVATED（激活），INACTIVATED（未激活）

	TimingStatus *ShowTaskResponseTimingStatus `json:"timing_status,omitempty"`

	Timing *TaskTiming `json:"timing,omitempty"`

	Input *TaskInput `json:"input"`

	Output *TaskOutputForDisplay `json:"output"`

	ServiceConfig *TaskServiceConfig `json:"service_config,omitempty"`

	HostingResult  *TaskHostingResultHostingResult `json:"hosting_result,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}

type ShowTaskResponseState struct {
	value string
}

type ShowTaskResponseStateEnum struct {
	PENDING       ShowTaskResponseState
	RECOVERING    ShowTaskResponseState
	STARTING      ShowTaskResponseState
	UPGRADING     ShowTaskResponseState
	CREATE_FAILED ShowTaskResponseState
	START_FAILED  ShowTaskResponseState
	RUNNING       ShowTaskResponseState
	STOPPING      ShowTaskResponseState
	STOPPED       ShowTaskResponseState
	ABNORMAL      ShowTaskResponseState
	SUCCEEDED     ShowTaskResponseState
	FAILED        ShowTaskResponseState
	DELETING      ShowTaskResponseState
	FREEZING      ShowTaskResponseState
	FROZEN        ShowTaskResponseState
}

func GetShowTaskResponseStateEnum() ShowTaskResponseStateEnum {
	return ShowTaskResponseStateEnum{
		PENDING: ShowTaskResponseState{
			value: "PENDING",
		},
		RECOVERING: ShowTaskResponseState{
			value: "RECOVERING",
		},
		STARTING: ShowTaskResponseState{
			value: "STARTING",
		},
		UPGRADING: ShowTaskResponseState{
			value: "UPGRADING",
		},
		CREATE_FAILED: ShowTaskResponseState{
			value: "CREATE_FAILED",
		},
		START_FAILED: ShowTaskResponseState{
			value: "START_FAILED",
		},
		RUNNING: ShowTaskResponseState{
			value: "RUNNING",
		},
		STOPPING: ShowTaskResponseState{
			value: "STOPPING",
		},
		STOPPED: ShowTaskResponseState{
			value: "STOPPED",
		},
		ABNORMAL: ShowTaskResponseState{
			value: "ABNORMAL",
		},
		SUCCEEDED: ShowTaskResponseState{
			value: "SUCCEEDED",
		},
		FAILED: ShowTaskResponseState{
			value: "FAILED",
		},
		DELETING: ShowTaskResponseState{
			value: "DELETING",
		},
		FREEZING: ShowTaskResponseState{
			value: "FREEZING",
		},
		FROZEN: ShowTaskResponseState{
			value: "FROZEN",
		},
	}
}

func (c ShowTaskResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseState) UnmarshalJSON(b []byte) error {
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

type ShowTaskResponseTimingStatus struct {
	value string
}

type ShowTaskResponseTimingStatusEnum struct {
	ACTIVATED   ShowTaskResponseTimingStatus
	INACTIVATED ShowTaskResponseTimingStatus
}

func GetShowTaskResponseTimingStatusEnum() ShowTaskResponseTimingStatusEnum {
	return ShowTaskResponseTimingStatusEnum{
		ACTIVATED: ShowTaskResponseTimingStatus{
			value: "ACTIVATED",
		},
		INACTIVATED: ShowTaskResponseTimingStatus{
			value: "INACTIVATED",
		},
	}
}

func (c ShowTaskResponseTimingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskResponseTimingStatus) UnmarshalJSON(b []byte) error {
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
