package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListOpLogsRequest struct {

	// 任务结束时间，格式为%YYYY-%mm-%ddT%HH:%MM:%SSZ，例如2018-02-01T12:00:00Z
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 每页显示的条目数量，正整数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移值，正整数
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 任务类型
	OperationType *ListOpLogsRequestOperationType `json:"operation_type,omitempty" xml:"operation_type"`

	// 备份提供商ID
	ProviderId *string `json:"provider_id,omitempty" xml:"provider_id"`

	// 该任务操作的资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 该任务操作的资源名称
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 任务开始时间，格式为%YYYY-%mm-%ddT%HH:%MM:%SSZ，例如2018-01-01T12:00:00Z
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 任务状态
	Status *ListOpLogsRequestStatus `json:"status,omitempty" xml:"status"`

	// 存储库ID,该任务操作的资源所属绑定的存储库。
	VaultId *string `json:"vault_id,omitempty" xml:"vault_id"`

	// 存储库名称，该任务操作资源所绑定的存储库名称。
	VaultName *string `json:"vault_name,omitempty" xml:"vault_name"`

	// 企业项目id或all_granted_eps，all_granted_eps表示查询用户有权限的所有企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`
}

func (o ListOpLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpLogsRequest struct{}"
	}

	return strings.Join([]string{"ListOpLogsRequest", string(data)}, " ")
}

type ListOpLogsRequestOperationType struct {
	value string
}

type ListOpLogsRequestOperationTypeEnum struct {
	BACKUP          ListOpLogsRequestOperationType
	COPY            ListOpLogsRequestOperationType
	REPLICATION     ListOpLogsRequestOperationType
	DELETE          ListOpLogsRequestOperationType
	RESTORE         ListOpLogsRequestOperationType
	VAULT_DELETE    ListOpLogsRequestOperationType
	REMOVE_RESOURCE ListOpLogsRequestOperationType
	SYNC            ListOpLogsRequestOperationType
}

func GetListOpLogsRequestOperationTypeEnum() ListOpLogsRequestOperationTypeEnum {
	return ListOpLogsRequestOperationTypeEnum{
		BACKUP: ListOpLogsRequestOperationType{
			value: "backup",
		},
		COPY: ListOpLogsRequestOperationType{
			value: "copy",
		},
		REPLICATION: ListOpLogsRequestOperationType{
			value: "replication",
		},
		DELETE: ListOpLogsRequestOperationType{
			value: "delete",
		},
		RESTORE: ListOpLogsRequestOperationType{
			value: "restore",
		},
		VAULT_DELETE: ListOpLogsRequestOperationType{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: ListOpLogsRequestOperationType{
			value: "remove_resource",
		},
		SYNC: ListOpLogsRequestOperationType{
			value: "sync",
		},
	}
}

func (c ListOpLogsRequestOperationType) Value() string {
	return c.value
}

func (c ListOpLogsRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpLogsRequestOperationType) UnmarshalJSON(b []byte) error {
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

type ListOpLogsRequestStatus struct {
	value string
}

type ListOpLogsRequestStatusEnum struct {
	SUCCESS ListOpLogsRequestStatus
	SKIPPED ListOpLogsRequestStatus
	FAILED  ListOpLogsRequestStatus
	RUNNING ListOpLogsRequestStatus
	TIMEOUT ListOpLogsRequestStatus
	WAITING ListOpLogsRequestStatus
}

func GetListOpLogsRequestStatusEnum() ListOpLogsRequestStatusEnum {
	return ListOpLogsRequestStatusEnum{
		SUCCESS: ListOpLogsRequestStatus{
			value: "success",
		},
		SKIPPED: ListOpLogsRequestStatus{
			value: "skipped",
		},
		FAILED: ListOpLogsRequestStatus{
			value: "failed",
		},
		RUNNING: ListOpLogsRequestStatus{
			value: "running",
		},
		TIMEOUT: ListOpLogsRequestStatus{
			value: "timeout",
		},
		WAITING: ListOpLogsRequestStatus{
			value: "waiting",
		},
	}
}

func (c ListOpLogsRequestStatus) Value() string {
	return c.value
}

func (c ListOpLogsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpLogsRequestStatus) UnmarshalJSON(b []byte) error {
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
