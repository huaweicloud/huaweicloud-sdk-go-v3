package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTaskStatusSummaryRequest Request Object
type ShowTaskStatusSummaryRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// Region ID
	RegionId *string `json:"region_id,omitempty"`

	// 任务类型，取值范围：backup,replication,restore,delete,vault_delete,remove_resource
	TaskType *ShowTaskStatusSummaryRequestTaskType `json:"task_type,omitempty"`

	// 资源类型，取值范围：Server,Volume,Sfs-Turbo,Workspace,MySQL,PostgreSQL,SQLServer,MariaDB,GaussDB
	ResourceType *ShowTaskStatusSummaryRequestResourceType `json:"resource_type,omitempty"`

	// 查询范围的起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围的结束时间，例如：2025-03-21T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowTaskStatusSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskStatusSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskStatusSummaryRequest", string(data)}, " ")
}

type ShowTaskStatusSummaryRequestTaskType struct {
	value string
}

type ShowTaskStatusSummaryRequestTaskTypeEnum struct {
	BACKUP          ShowTaskStatusSummaryRequestTaskType
	REPLICATION     ShowTaskStatusSummaryRequestTaskType
	RESTORE         ShowTaskStatusSummaryRequestTaskType
	DELETE          ShowTaskStatusSummaryRequestTaskType
	VAULT_DELETE    ShowTaskStatusSummaryRequestTaskType
	REMOVE_RESOURCE ShowTaskStatusSummaryRequestTaskType
}

func GetShowTaskStatusSummaryRequestTaskTypeEnum() ShowTaskStatusSummaryRequestTaskTypeEnum {
	return ShowTaskStatusSummaryRequestTaskTypeEnum{
		BACKUP: ShowTaskStatusSummaryRequestTaskType{
			value: "backup",
		},
		REPLICATION: ShowTaskStatusSummaryRequestTaskType{
			value: "replication",
		},
		RESTORE: ShowTaskStatusSummaryRequestTaskType{
			value: "restore",
		},
		DELETE: ShowTaskStatusSummaryRequestTaskType{
			value: "delete",
		},
		VAULT_DELETE: ShowTaskStatusSummaryRequestTaskType{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: ShowTaskStatusSummaryRequestTaskType{
			value: "remove_resource",
		},
	}
}

func (c ShowTaskStatusSummaryRequestTaskType) Value() string {
	return c.value
}

func (c ShowTaskStatusSummaryRequestTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskStatusSummaryRequestTaskType) UnmarshalJSON(b []byte) error {
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

type ShowTaskStatusSummaryRequestResourceType struct {
	value string
}

type ShowTaskStatusSummaryRequestResourceTypeEnum struct {
	SERVER      ShowTaskStatusSummaryRequestResourceType
	VOLUME      ShowTaskStatusSummaryRequestResourceType
	SFS_TURBO   ShowTaskStatusSummaryRequestResourceType
	WORKSPACE   ShowTaskStatusSummaryRequestResourceType
	MY_SQL      ShowTaskStatusSummaryRequestResourceType
	POSTGRE_SQL ShowTaskStatusSummaryRequestResourceType
	SQL_SERVER  ShowTaskStatusSummaryRequestResourceType
	MARIA_DB    ShowTaskStatusSummaryRequestResourceType
	GAUSS_DB    ShowTaskStatusSummaryRequestResourceType
}

func GetShowTaskStatusSummaryRequestResourceTypeEnum() ShowTaskStatusSummaryRequestResourceTypeEnum {
	return ShowTaskStatusSummaryRequestResourceTypeEnum{
		SERVER: ShowTaskStatusSummaryRequestResourceType{
			value: "Server",
		},
		VOLUME: ShowTaskStatusSummaryRequestResourceType{
			value: "Volume",
		},
		SFS_TURBO: ShowTaskStatusSummaryRequestResourceType{
			value: "Sfs-Turbo",
		},
		WORKSPACE: ShowTaskStatusSummaryRequestResourceType{
			value: "Workspace",
		},
		MY_SQL: ShowTaskStatusSummaryRequestResourceType{
			value: "MySQL",
		},
		POSTGRE_SQL: ShowTaskStatusSummaryRequestResourceType{
			value: "PostgreSQL",
		},
		SQL_SERVER: ShowTaskStatusSummaryRequestResourceType{
			value: "SQLServer",
		},
		MARIA_DB: ShowTaskStatusSummaryRequestResourceType{
			value: "MariaDB",
		},
		GAUSS_DB: ShowTaskStatusSummaryRequestResourceType{
			value: "GaussDB",
		},
	}
}

func (c ShowTaskStatusSummaryRequestResourceType) Value() string {
	return c.value
}

func (c ShowTaskStatusSummaryRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskStatusSummaryRequestResourceType) UnmarshalJSON(b []byte) error {
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
