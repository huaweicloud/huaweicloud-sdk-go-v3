package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityMemberSyncTasksRequest Request Object
type ListSecurityMemberSyncTasksRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// 集群类型 * MRS数据源 * DWS数据源
	ClusterType *ListSecurityMemberSyncTasksRequestClusterType `json:"cluster_type,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 同步状态 * UNKNOWN 未知 * NOT_SYNC 未同步 * SYNCING 同步中 * SYNC_SUCCESS 同步成功 * SYNC_FAIL 同步失败
	SyncStatus *ListSecurityMemberSyncTasksRequestSyncStatus `json:"sync_status,omitempty"`

	// 用户同步任务调度状态 * NOT_SCHEDULE 未启用调度 * SCHEDULING 调度中
	ScheduleStatus *ListSecurityMemberSyncTasksRequestScheduleStatus `json:"schedule_status,omitempty"`

	// 排序字段 * CLUSTER_NAME  按照集群名称排序 * CREATE_TIME   按照创建时间排序 * UPDATE_TIME   按照更新时间排序 * SYNC_TIME     按照同步时间排序
	OrderBy *ListSecurityMemberSyncTasksRequestOrderBy `json:"order_by,omitempty"`

	// 是否升序（仅指定排序参数时有效）。
	OrderByAsc *bool `json:"order_by_asc,omitempty"`
}

func (o ListSecurityMemberSyncTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityMemberSyncTasksRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityMemberSyncTasksRequest", string(data)}, " ")
}

type ListSecurityMemberSyncTasksRequestClusterType struct {
	value string
}

type ListSecurityMemberSyncTasksRequestClusterTypeEnum struct {
	MRS ListSecurityMemberSyncTasksRequestClusterType
	DWS ListSecurityMemberSyncTasksRequestClusterType
}

func GetListSecurityMemberSyncTasksRequestClusterTypeEnum() ListSecurityMemberSyncTasksRequestClusterTypeEnum {
	return ListSecurityMemberSyncTasksRequestClusterTypeEnum{
		MRS: ListSecurityMemberSyncTasksRequestClusterType{
			value: "MRS",
		},
		DWS: ListSecurityMemberSyncTasksRequestClusterType{
			value: "DWS",
		},
	}
}

func (c ListSecurityMemberSyncTasksRequestClusterType) Value() string {
	return c.value
}

func (c ListSecurityMemberSyncTasksRequestClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberSyncTasksRequestClusterType) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberSyncTasksRequestSyncStatus struct {
	value string
}

type ListSecurityMemberSyncTasksRequestSyncStatusEnum struct {
	UNKNOWN      ListSecurityMemberSyncTasksRequestSyncStatus
	NOT_SYNC     ListSecurityMemberSyncTasksRequestSyncStatus
	SYNCING      ListSecurityMemberSyncTasksRequestSyncStatus
	SYNC_SUCCESS ListSecurityMemberSyncTasksRequestSyncStatus
	SYNC_FAIL    ListSecurityMemberSyncTasksRequestSyncStatus
}

func GetListSecurityMemberSyncTasksRequestSyncStatusEnum() ListSecurityMemberSyncTasksRequestSyncStatusEnum {
	return ListSecurityMemberSyncTasksRequestSyncStatusEnum{
		UNKNOWN: ListSecurityMemberSyncTasksRequestSyncStatus{
			value: "UNKNOWN",
		},
		NOT_SYNC: ListSecurityMemberSyncTasksRequestSyncStatus{
			value: "NOT_SYNC",
		},
		SYNCING: ListSecurityMemberSyncTasksRequestSyncStatus{
			value: "SYNCING",
		},
		SYNC_SUCCESS: ListSecurityMemberSyncTasksRequestSyncStatus{
			value: "SYNC_SUCCESS",
		},
		SYNC_FAIL: ListSecurityMemberSyncTasksRequestSyncStatus{
			value: "SYNC_FAIL",
		},
	}
}

func (c ListSecurityMemberSyncTasksRequestSyncStatus) Value() string {
	return c.value
}

func (c ListSecurityMemberSyncTasksRequestSyncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberSyncTasksRequestSyncStatus) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberSyncTasksRequestScheduleStatus struct {
	value string
}

type ListSecurityMemberSyncTasksRequestScheduleStatusEnum struct {
	NOT_SCHEDULE ListSecurityMemberSyncTasksRequestScheduleStatus
	SCHEDULING   ListSecurityMemberSyncTasksRequestScheduleStatus
}

func GetListSecurityMemberSyncTasksRequestScheduleStatusEnum() ListSecurityMemberSyncTasksRequestScheduleStatusEnum {
	return ListSecurityMemberSyncTasksRequestScheduleStatusEnum{
		NOT_SCHEDULE: ListSecurityMemberSyncTasksRequestScheduleStatus{
			value: "NOT_SCHEDULE",
		},
		SCHEDULING: ListSecurityMemberSyncTasksRequestScheduleStatus{
			value: "SCHEDULING",
		},
	}
}

func (c ListSecurityMemberSyncTasksRequestScheduleStatus) Value() string {
	return c.value
}

func (c ListSecurityMemberSyncTasksRequestScheduleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberSyncTasksRequestScheduleStatus) UnmarshalJSON(b []byte) error {
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

type ListSecurityMemberSyncTasksRequestOrderBy struct {
	value string
}

type ListSecurityMemberSyncTasksRequestOrderByEnum struct {
	CLUSTER_NAME ListSecurityMemberSyncTasksRequestOrderBy
	CREATE_TIME  ListSecurityMemberSyncTasksRequestOrderBy
	UPDATE_TIME  ListSecurityMemberSyncTasksRequestOrderBy
	SYNC_TIME    ListSecurityMemberSyncTasksRequestOrderBy
}

func GetListSecurityMemberSyncTasksRequestOrderByEnum() ListSecurityMemberSyncTasksRequestOrderByEnum {
	return ListSecurityMemberSyncTasksRequestOrderByEnum{
		CLUSTER_NAME: ListSecurityMemberSyncTasksRequestOrderBy{
			value: "CLUSTER_NAME",
		},
		CREATE_TIME: ListSecurityMemberSyncTasksRequestOrderBy{
			value: "CREATE_TIME",
		},
		UPDATE_TIME: ListSecurityMemberSyncTasksRequestOrderBy{
			value: "UPDATE_TIME",
		},
		SYNC_TIME: ListSecurityMemberSyncTasksRequestOrderBy{
			value: "SYNC_TIME",
		},
	}
}

func (c ListSecurityMemberSyncTasksRequestOrderBy) Value() string {
	return c.value
}

func (c ListSecurityMemberSyncTasksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityMemberSyncTasksRequestOrderBy) UnmarshalJSON(b []byte) error {
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
