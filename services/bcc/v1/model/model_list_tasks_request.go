package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// RegionID
	RegionId *string `json:"region_id,omitempty"`

	// 任务状态，取值范围：success,failed,running,skipped,timeout
	TaskStatus *ListTasksRequestTaskStatus `json:"task_status,omitempty"`

	// 任务类型，取值范围：backup,replication,restore,delete,vault_delete,remove_resource
	TaskType *ListTasksRequestTaskType `json:"task_type,omitempty"`

	// 资源类型，取值范围：Server,Volume,Sfs-Turbo,Workspace,MySQL,PostgreSQL,SQLServer,MariaDB,GaussDB
	ResourceType *ListTasksRequestResourceType `json:"resource_type,omitempty"`

	// 查询范围起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围结束时间，例如：2025-03-21T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`

	// 单次查询的条数限制,取值范围(0,100]，默认值为10，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}

type ListTasksRequestTaskStatus struct {
	value string
}

type ListTasksRequestTaskStatusEnum struct {
	SUCCESS ListTasksRequestTaskStatus
	FAILED  ListTasksRequestTaskStatus
	RUNNING ListTasksRequestTaskStatus
	SKIPPED ListTasksRequestTaskStatus
	TIMEOUT ListTasksRequestTaskStatus
}

func GetListTasksRequestTaskStatusEnum() ListTasksRequestTaskStatusEnum {
	return ListTasksRequestTaskStatusEnum{
		SUCCESS: ListTasksRequestTaskStatus{
			value: "success",
		},
		FAILED: ListTasksRequestTaskStatus{
			value: "failed",
		},
		RUNNING: ListTasksRequestTaskStatus{
			value: "running",
		},
		SKIPPED: ListTasksRequestTaskStatus{
			value: "skipped",
		},
		TIMEOUT: ListTasksRequestTaskStatus{
			value: "timeout",
		},
	}
}

func (c ListTasksRequestTaskStatus) Value() string {
	return c.value
}

func (c ListTasksRequestTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestTaskStatus) UnmarshalJSON(b []byte) error {
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

type ListTasksRequestTaskType struct {
	value string
}

type ListTasksRequestTaskTypeEnum struct {
	BACKUP          ListTasksRequestTaskType
	REPLICATION     ListTasksRequestTaskType
	RESTORE         ListTasksRequestTaskType
	DELETE          ListTasksRequestTaskType
	VAULT_DELETE    ListTasksRequestTaskType
	REMOVE_RESOURCE ListTasksRequestTaskType
}

func GetListTasksRequestTaskTypeEnum() ListTasksRequestTaskTypeEnum {
	return ListTasksRequestTaskTypeEnum{
		BACKUP: ListTasksRequestTaskType{
			value: "backup",
		},
		REPLICATION: ListTasksRequestTaskType{
			value: "replication",
		},
		RESTORE: ListTasksRequestTaskType{
			value: "restore",
		},
		DELETE: ListTasksRequestTaskType{
			value: "delete",
		},
		VAULT_DELETE: ListTasksRequestTaskType{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: ListTasksRequestTaskType{
			value: "remove_resource",
		},
	}
}

func (c ListTasksRequestTaskType) Value() string {
	return c.value
}

func (c ListTasksRequestTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestTaskType) UnmarshalJSON(b []byte) error {
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

type ListTasksRequestResourceType struct {
	value string
}

type ListTasksRequestResourceTypeEnum struct {
	SERVER      ListTasksRequestResourceType
	VOLUME      ListTasksRequestResourceType
	SFS_TURBO   ListTasksRequestResourceType
	WORKSPACE   ListTasksRequestResourceType
	MY_SQL      ListTasksRequestResourceType
	POSTGRE_SQL ListTasksRequestResourceType
	SQL_SERVER  ListTasksRequestResourceType
	MARIA_DB    ListTasksRequestResourceType
	GAUSS_DB    ListTasksRequestResourceType
}

func GetListTasksRequestResourceTypeEnum() ListTasksRequestResourceTypeEnum {
	return ListTasksRequestResourceTypeEnum{
		SERVER: ListTasksRequestResourceType{
			value: "Server",
		},
		VOLUME: ListTasksRequestResourceType{
			value: "Volume",
		},
		SFS_TURBO: ListTasksRequestResourceType{
			value: "Sfs-Turbo",
		},
		WORKSPACE: ListTasksRequestResourceType{
			value: "Workspace",
		},
		MY_SQL: ListTasksRequestResourceType{
			value: "MySQL",
		},
		POSTGRE_SQL: ListTasksRequestResourceType{
			value: "PostgreSQL",
		},
		SQL_SERVER: ListTasksRequestResourceType{
			value: "SQLServer",
		},
		MARIA_DB: ListTasksRequestResourceType{
			value: "MariaDB",
		},
		GAUSS_DB: ListTasksRequestResourceType{
			value: "GaussDB",
		},
	}
}

func (c ListTasksRequestResourceType) Value() string {
	return c.value
}

func (c ListTasksRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTasksRequestResourceType) UnmarshalJSON(b []byte) error {
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
