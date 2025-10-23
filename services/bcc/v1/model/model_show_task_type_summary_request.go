package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTaskTypeSummaryRequest Request Object
type ShowTaskTypeSummaryRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// Region ID
	RegionId *string `json:"region_id,omitempty"`

	// 任务状态，取值范围success,failed,running,skipped,timeout
	TaskStatus *ShowTaskTypeSummaryRequestTaskStatus `json:"task_status,omitempty"`

	// 资源类型，取值范围：Server,Volume,Sfs-Turbo,Workspace,MySQL,PostgreSQL,SQLServer,MariaDB,GaussDB
	ResourceType *ShowTaskTypeSummaryRequestResourceType `json:"resource_type,omitempty"`

	// 查询范围的起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围的结束时间，例如：2025-03-21T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowTaskTypeSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskTypeSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskTypeSummaryRequest", string(data)}, " ")
}

type ShowTaskTypeSummaryRequestTaskStatus struct {
	value string
}

type ShowTaskTypeSummaryRequestTaskStatusEnum struct {
	SUCCESS ShowTaskTypeSummaryRequestTaskStatus
	FAILED  ShowTaskTypeSummaryRequestTaskStatus
	RUNNING ShowTaskTypeSummaryRequestTaskStatus
	SKIPPED ShowTaskTypeSummaryRequestTaskStatus
	TIMEOUT ShowTaskTypeSummaryRequestTaskStatus
}

func GetShowTaskTypeSummaryRequestTaskStatusEnum() ShowTaskTypeSummaryRequestTaskStatusEnum {
	return ShowTaskTypeSummaryRequestTaskStatusEnum{
		SUCCESS: ShowTaskTypeSummaryRequestTaskStatus{
			value: "success",
		},
		FAILED: ShowTaskTypeSummaryRequestTaskStatus{
			value: "failed",
		},
		RUNNING: ShowTaskTypeSummaryRequestTaskStatus{
			value: "running",
		},
		SKIPPED: ShowTaskTypeSummaryRequestTaskStatus{
			value: "skipped",
		},
		TIMEOUT: ShowTaskTypeSummaryRequestTaskStatus{
			value: "timeout",
		},
	}
}

func (c ShowTaskTypeSummaryRequestTaskStatus) Value() string {
	return c.value
}

func (c ShowTaskTypeSummaryRequestTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskTypeSummaryRequestTaskStatus) UnmarshalJSON(b []byte) error {
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

type ShowTaskTypeSummaryRequestResourceType struct {
	value string
}

type ShowTaskTypeSummaryRequestResourceTypeEnum struct {
	SERVER      ShowTaskTypeSummaryRequestResourceType
	VOLUME      ShowTaskTypeSummaryRequestResourceType
	SFS_TURBO   ShowTaskTypeSummaryRequestResourceType
	WORKSPACE   ShowTaskTypeSummaryRequestResourceType
	MY_SQL      ShowTaskTypeSummaryRequestResourceType
	POSTGRE_SQL ShowTaskTypeSummaryRequestResourceType
	SQL_SERVER  ShowTaskTypeSummaryRequestResourceType
	MARIA_DB    ShowTaskTypeSummaryRequestResourceType
	GAUSS_DB    ShowTaskTypeSummaryRequestResourceType
}

func GetShowTaskTypeSummaryRequestResourceTypeEnum() ShowTaskTypeSummaryRequestResourceTypeEnum {
	return ShowTaskTypeSummaryRequestResourceTypeEnum{
		SERVER: ShowTaskTypeSummaryRequestResourceType{
			value: "Server",
		},
		VOLUME: ShowTaskTypeSummaryRequestResourceType{
			value: "Volume",
		},
		SFS_TURBO: ShowTaskTypeSummaryRequestResourceType{
			value: "Sfs-Turbo",
		},
		WORKSPACE: ShowTaskTypeSummaryRequestResourceType{
			value: "Workspace",
		},
		MY_SQL: ShowTaskTypeSummaryRequestResourceType{
			value: "MySQL",
		},
		POSTGRE_SQL: ShowTaskTypeSummaryRequestResourceType{
			value: "PostgreSQL",
		},
		SQL_SERVER: ShowTaskTypeSummaryRequestResourceType{
			value: "SQLServer",
		},
		MARIA_DB: ShowTaskTypeSummaryRequestResourceType{
			value: "MariaDB",
		},
		GAUSS_DB: ShowTaskTypeSummaryRequestResourceType{
			value: "GaussDB",
		},
	}
}

func (c ShowTaskTypeSummaryRequestResourceType) Value() string {
	return c.value
}

func (c ShowTaskTypeSummaryRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskTypeSummaryRequestResourceType) UnmarshalJSON(b []byte) error {
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
