package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MemberSyncTask struct {

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
	ClusterType *MemberSyncTaskClusterType `json:"cluster_type,omitempty"`

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
	SchedulePeriod *MemberSyncTaskSchedulePeriod `json:"schedule_period,omitempty"`

	// 调度间隔。
	ScheduleInterval *int32 `json:"schedule_interval,omitempty"`

	// 调度状态 * NOT_SCHEDULE  未启用任务调度 * SCHEDULING    任务调度中
	ScheduleStatus *MemberSyncTaskScheduleStatus `json:"schedule_status,omitempty"`

	// 同步状态 * UNKNOWN 未知 * NOT_SYNC 未同步 * SYNCING 同步中 * SYNC_SUCCESS 同步成功 * SYNC_FAIL 同步失败
	SyncStatus *MemberSyncTaskSyncStatus `json:"sync_status,omitempty"`

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
	UpdateUser *string `json:"update_user,omitempty"`
}

func (o MemberSyncTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberSyncTask struct{}"
	}

	return strings.Join([]string{"MemberSyncTask", string(data)}, " ")
}

type MemberSyncTaskClusterType struct {
	value string
}

type MemberSyncTaskClusterTypeEnum struct {
	MRS MemberSyncTaskClusterType
	DWS MemberSyncTaskClusterType
}

func GetMemberSyncTaskClusterTypeEnum() MemberSyncTaskClusterTypeEnum {
	return MemberSyncTaskClusterTypeEnum{
		MRS: MemberSyncTaskClusterType{
			value: "MRS",
		},
		DWS: MemberSyncTaskClusterType{
			value: "DWS",
		},
	}
}

func (c MemberSyncTaskClusterType) Value() string {
	return c.value
}

func (c MemberSyncTaskClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberSyncTaskClusterType) UnmarshalJSON(b []byte) error {
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

type MemberSyncTaskSchedulePeriod struct {
	value string
}

type MemberSyncTaskSchedulePeriodEnum struct {
	MINUTE MemberSyncTaskSchedulePeriod
	HOUR   MemberSyncTaskSchedulePeriod
}

func GetMemberSyncTaskSchedulePeriodEnum() MemberSyncTaskSchedulePeriodEnum {
	return MemberSyncTaskSchedulePeriodEnum{
		MINUTE: MemberSyncTaskSchedulePeriod{
			value: "MINUTE",
		},
		HOUR: MemberSyncTaskSchedulePeriod{
			value: "HOUR",
		},
	}
}

func (c MemberSyncTaskSchedulePeriod) Value() string {
	return c.value
}

func (c MemberSyncTaskSchedulePeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberSyncTaskSchedulePeriod) UnmarshalJSON(b []byte) error {
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

type MemberSyncTaskScheduleStatus struct {
	value string
}

type MemberSyncTaskScheduleStatusEnum struct {
	NOT_SCHEDULE MemberSyncTaskScheduleStatus
	SCHEDULING   MemberSyncTaskScheduleStatus
}

func GetMemberSyncTaskScheduleStatusEnum() MemberSyncTaskScheduleStatusEnum {
	return MemberSyncTaskScheduleStatusEnum{
		NOT_SCHEDULE: MemberSyncTaskScheduleStatus{
			value: "NOT_SCHEDULE",
		},
		SCHEDULING: MemberSyncTaskScheduleStatus{
			value: "SCHEDULING",
		},
	}
}

func (c MemberSyncTaskScheduleStatus) Value() string {
	return c.value
}

func (c MemberSyncTaskScheduleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberSyncTaskScheduleStatus) UnmarshalJSON(b []byte) error {
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

type MemberSyncTaskSyncStatus struct {
	value string
}

type MemberSyncTaskSyncStatusEnum struct {
	UNKNOWN      MemberSyncTaskSyncStatus
	NOT_SYNC     MemberSyncTaskSyncStatus
	SYNCING      MemberSyncTaskSyncStatus
	SYNC_SUCCESS MemberSyncTaskSyncStatus
	SYNC_FAIL    MemberSyncTaskSyncStatus
}

func GetMemberSyncTaskSyncStatusEnum() MemberSyncTaskSyncStatusEnum {
	return MemberSyncTaskSyncStatusEnum{
		UNKNOWN: MemberSyncTaskSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: MemberSyncTaskSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: MemberSyncTaskSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: MemberSyncTaskSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: MemberSyncTaskSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c MemberSyncTaskSyncStatus) Value() string {
	return c.value
}

func (c MemberSyncTaskSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberSyncTaskSyncStatus) UnmarshalJSON(b []byte) error {
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
