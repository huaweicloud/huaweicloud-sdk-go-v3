package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRealTimeJobDetailsResponse Response Object
type ShowRealTimeJobDetailsResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业状态。 - EXCEPTION：异常 - STOPPING：停止中 - SUBMITTING：提交中 - RUNNING：运行中 - STOPPED：已停止 - SUCCESS：成功
	State *ShowRealTimeJobDetailsResponseState `json:"state,omitempty"`

	// 作业迁移类型。 - INCREMENTAL_DATA：增量数据 - HISTORY_DATA：历史数据
	MigrationType *ShowRealTimeJobDetailsResponseMigrationType `json:"migration_type,omitempty"`

	// INCREMENTAL_DATA作业启动的时间位点。
	StartupTimestamp *string `json:"startup_timestamp,omitempty"`

	// 运行作业时的引擎版本。
	JobEngineVersion *string `json:"job_engine_version,omitempty"`

	// 作业关联资源组的引擎版本。
	ClusterEngineVersion *string `json:"cluster_engine_version,omitempty"`

	// 资源组类型。
	ClusterType *string `json:"cluster_type,omitempty"`

	// MRS Flink作业trackingUrl。
	TrackingUrl *string `json:"tracking_url,omitempty"`

	// 作业指标信息。
	MetricInfo *string `json:"metric_info,omitempty"`

	// 任务详情列表。
	TaskDetails    *[]TaskDetailInfo `json:"task_details,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowRealTimeJobDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealTimeJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowRealTimeJobDetailsResponse", string(data)}, " ")
}

type ShowRealTimeJobDetailsResponseState struct {
	value string
}

type ShowRealTimeJobDetailsResponseStateEnum struct {
	EXCEPTION  ShowRealTimeJobDetailsResponseState
	STOPPING   ShowRealTimeJobDetailsResponseState
	SUBMITTING ShowRealTimeJobDetailsResponseState
	RUNNING    ShowRealTimeJobDetailsResponseState
	STOPPED    ShowRealTimeJobDetailsResponseState
	SUCCESS    ShowRealTimeJobDetailsResponseState
}

func GetShowRealTimeJobDetailsResponseStateEnum() ShowRealTimeJobDetailsResponseStateEnum {
	return ShowRealTimeJobDetailsResponseStateEnum{
		EXCEPTION: ShowRealTimeJobDetailsResponseState{
			value: "EXCEPTION",
		},
		STOPPING: ShowRealTimeJobDetailsResponseState{
			value: "STOPPING",
		},
		SUBMITTING: ShowRealTimeJobDetailsResponseState{
			value: "SUBMITTING",
		},
		RUNNING: ShowRealTimeJobDetailsResponseState{
			value: "RUNNING",
		},
		STOPPED: ShowRealTimeJobDetailsResponseState{
			value: "STOPPED",
		},
		SUCCESS: ShowRealTimeJobDetailsResponseState{
			value: "SUCCESS",
		},
	}
}

func (c ShowRealTimeJobDetailsResponseState) Value() string {
	return c.value
}

func (c ShowRealTimeJobDetailsResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRealTimeJobDetailsResponseState) UnmarshalJSON(b []byte) error {
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

type ShowRealTimeJobDetailsResponseMigrationType struct {
	value string
}

type ShowRealTimeJobDetailsResponseMigrationTypeEnum struct {
	INCREMENTAL_DATA ShowRealTimeJobDetailsResponseMigrationType
	HISTORY_DATA     ShowRealTimeJobDetailsResponseMigrationType
}

func GetShowRealTimeJobDetailsResponseMigrationTypeEnum() ShowRealTimeJobDetailsResponseMigrationTypeEnum {
	return ShowRealTimeJobDetailsResponseMigrationTypeEnum{
		INCREMENTAL_DATA: ShowRealTimeJobDetailsResponseMigrationType{
			value: "INCREMENTAL_DATA",
		},
		HISTORY_DATA: ShowRealTimeJobDetailsResponseMigrationType{
			value: "HISTORY_DATA",
		},
	}
}

func (c ShowRealTimeJobDetailsResponseMigrationType) Value() string {
	return c.value
}

func (c ShowRealTimeJobDetailsResponseMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRealTimeJobDetailsResponseMigrationType) UnmarshalJSON(b []byte) error {
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
