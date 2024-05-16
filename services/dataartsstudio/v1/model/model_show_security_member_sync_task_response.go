package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityMemberSyncTaskResponse Response Object
type ShowSecurityMemberSyncTaskResponse struct {

	// 用户同步任务id。
	Id *string `json:"id,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 数据连接工作空间ID。
	DataConnectionWorkspace *string `json:"data_connection_workspace,omitempty"`

	// 集群类型 * MRS集群 * DWS集群
	ClusterType *ShowSecurityMemberSyncTaskResponseClusterType `json:"cluster_type,omitempty"`

	// 数据连接id。
	DataConnectionId *string `json:"data_connection_id,omitempty"`

	// 数据连接名称。
	DataConnectionName *string `json:"data_connection_name,omitempty"`

	// 集群id。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 调度开始时间, 单位为小时, 0~23。
	ScheduleStartHour *int32 `json:"schedule_start_hour,omitempty"`

	// 调度结束时间, 单位为小时, 0~23。
	ScheduleEndHour *int32 `json:"schedule_end_hour,omitempty"`

	// 调度周期 * MINUTE  分钟为单位调度 * HOUR    小时为单位调度
	SchedulePeriod *ShowSecurityMemberSyncTaskResponseSchedulePeriod `json:"schedule_period,omitempty"`

	// 调度间隔。
	ScheduleInterval *int32 `json:"schedule_interval,omitempty"`

	// 调度状态 * NOT_SCHEDULE  未启用任务调度 * SCHEDULING    任务调度中
	ScheduleStatus *ShowSecurityMemberSyncTaskResponseScheduleStatus `json:"schedule_status,omitempty"`

	// 同步状态 * UNKNOWN 未知 * NOT_SYNC 未同步 * SYNCING 同步中 * SYNC_SUCCESS 同步成功 * SYNC_FAIL 同步失败
	SyncStatus *ShowSecurityMemberSyncTaskResponseSyncStatus `json:"sync_status,omitempty"`

	// 同步日志。
	SyncMsg *string `json:"sync_msg,omitempty"`

	// 同步时间。
	SyncTime *int64 `json:"sync_time,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创建者。
	CreateUser *string `json:"create_user,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 更新者。
	UpdateUser     *string `json:"update_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecurityMemberSyncTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityMemberSyncTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityMemberSyncTaskResponse", string(data)}, " ")
}

type ShowSecurityMemberSyncTaskResponseClusterType struct {
	value string
}

type ShowSecurityMemberSyncTaskResponseClusterTypeEnum struct {
	MRS ShowSecurityMemberSyncTaskResponseClusterType
	DWS ShowSecurityMemberSyncTaskResponseClusterType
}

func GetShowSecurityMemberSyncTaskResponseClusterTypeEnum() ShowSecurityMemberSyncTaskResponseClusterTypeEnum {
	return ShowSecurityMemberSyncTaskResponseClusterTypeEnum{
		MRS: ShowSecurityMemberSyncTaskResponseClusterType{
			value: "MRS",
		},
		DWS: ShowSecurityMemberSyncTaskResponseClusterType{
			value: "DWS",
		},
	}
}

func (c ShowSecurityMemberSyncTaskResponseClusterType) Value() string {
	return c.value
}

func (c ShowSecurityMemberSyncTaskResponseClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityMemberSyncTaskResponseClusterType) UnmarshalJSON(b []byte) error {
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

type ShowSecurityMemberSyncTaskResponseSchedulePeriod struct {
	value string
}

type ShowSecurityMemberSyncTaskResponseSchedulePeriodEnum struct {
	MINUTE ShowSecurityMemberSyncTaskResponseSchedulePeriod
	HOUR   ShowSecurityMemberSyncTaskResponseSchedulePeriod
}

func GetShowSecurityMemberSyncTaskResponseSchedulePeriodEnum() ShowSecurityMemberSyncTaskResponseSchedulePeriodEnum {
	return ShowSecurityMemberSyncTaskResponseSchedulePeriodEnum{
		MINUTE: ShowSecurityMemberSyncTaskResponseSchedulePeriod{
			value: "MINUTE",
		},
		HOUR: ShowSecurityMemberSyncTaskResponseSchedulePeriod{
			value: "HOUR",
		},
	}
}

func (c ShowSecurityMemberSyncTaskResponseSchedulePeriod) Value() string {
	return c.value
}

func (c ShowSecurityMemberSyncTaskResponseSchedulePeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityMemberSyncTaskResponseSchedulePeriod) UnmarshalJSON(b []byte) error {
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

type ShowSecurityMemberSyncTaskResponseScheduleStatus struct {
	value string
}

type ShowSecurityMemberSyncTaskResponseScheduleStatusEnum struct {
	NOT_SCHEDULE ShowSecurityMemberSyncTaskResponseScheduleStatus
	SCHEDULING   ShowSecurityMemberSyncTaskResponseScheduleStatus
}

func GetShowSecurityMemberSyncTaskResponseScheduleStatusEnum() ShowSecurityMemberSyncTaskResponseScheduleStatusEnum {
	return ShowSecurityMemberSyncTaskResponseScheduleStatusEnum{
		NOT_SCHEDULE: ShowSecurityMemberSyncTaskResponseScheduleStatus{
			value: "NOT_SCHEDULE",
		},
		SCHEDULING: ShowSecurityMemberSyncTaskResponseScheduleStatus{
			value: "SCHEDULING",
		},
	}
}

func (c ShowSecurityMemberSyncTaskResponseScheduleStatus) Value() string {
	return c.value
}

func (c ShowSecurityMemberSyncTaskResponseScheduleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityMemberSyncTaskResponseScheduleStatus) UnmarshalJSON(b []byte) error {
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

type ShowSecurityMemberSyncTaskResponseSyncStatus struct {
	value string
}

type ShowSecurityMemberSyncTaskResponseSyncStatusEnum struct {
	UNKNOWN      ShowSecurityMemberSyncTaskResponseSyncStatus
	NOT_SYNC     ShowSecurityMemberSyncTaskResponseSyncStatus
	SYNCING      ShowSecurityMemberSyncTaskResponseSyncStatus
	SYNC_SUCCESS ShowSecurityMemberSyncTaskResponseSyncStatus
	SYNC_FAIL    ShowSecurityMemberSyncTaskResponseSyncStatus
}

func GetShowSecurityMemberSyncTaskResponseSyncStatusEnum() ShowSecurityMemberSyncTaskResponseSyncStatusEnum {
	return ShowSecurityMemberSyncTaskResponseSyncStatusEnum{
		UNKNOWN: ShowSecurityMemberSyncTaskResponseSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: ShowSecurityMemberSyncTaskResponseSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: ShowSecurityMemberSyncTaskResponseSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: ShowSecurityMemberSyncTaskResponseSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: ShowSecurityMemberSyncTaskResponseSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c ShowSecurityMemberSyncTaskResponseSyncStatus) Value() string {
	return c.value
}

func (c ShowSecurityMemberSyncTaskResponseSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityMemberSyncTaskResponseSyncStatus) UnmarshalJSON(b []byte) error {
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
