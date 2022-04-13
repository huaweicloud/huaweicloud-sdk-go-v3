package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type OperationLog struct {
	// 备份记录id

	CheckpointId *string `json:"checkpoint_id,omitempty"`
	// 创建时间,例如: \"2020-02-23T01:00:32Z\"

	CreatedAt string `json:"created_at"`
	// 任务结束时间,例如: \"2020-02-23T01:00:32Z\"

	EndedAt *string `json:"ended_at,omitempty"`

	ErrorInfo *OpErrorInfo `json:"error_info"`

	ExtraInfo *OpExtraInfo `json:"extra_info,omitempty"`
	// 任务id

	Id string `json:"id"`
	// 任务类型

	OperationType *OperationLogOperationType `json:"operation_type,omitempty"`
	// 策略ID

	PolicyId *string `json:"policy_id,omitempty"`
	// 项目ID

	ProjectId string `json:"project_id"`
	// 备份提供商ID。用于区分备份对象。

	ProviderId *string `json:"provider_id,omitempty"`
	// 任务开始时间,例如: \"2020-02-23T01:00:32Z\"

	StartedAt string `json:"started_at"`
	// 任务状态

	Status OperationLogStatus `json:"status"`
	// 修改时间,例如: \"2020-02-23T01:00:32Z\"

	UpdatedAt string `json:"updated_at"`
	// 任务操作资源所属存储库ID

	VaultId *string `json:"vault_id,omitempty"`
	// 任务操作资源所属存储库名称

	VaultName *string `json:"vault_name,omitempty"`
}

func (o OperationLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationLog struct{}"
	}

	return strings.Join([]string{"OperationLog", string(data)}, " ")
}

type OperationLogOperationType struct {
	value string
}

type OperationLogOperationTypeEnum struct {
	BACKUP          OperationLogOperationType
	COPY            OperationLogOperationType
	REPLICATION     OperationLogOperationType
	RESTORE         OperationLogOperationType
	DELETE          OperationLogOperationType
	SYNC            OperationLogOperationType
	VAULT_DELETE    OperationLogOperationType
	REMOVE_RESOURCE OperationLogOperationType
}

func GetOperationLogOperationTypeEnum() OperationLogOperationTypeEnum {
	return OperationLogOperationTypeEnum{
		BACKUP: OperationLogOperationType{
			value: "backup",
		},
		COPY: OperationLogOperationType{
			value: "copy",
		},
		REPLICATION: OperationLogOperationType{
			value: "replication",
		},
		RESTORE: OperationLogOperationType{
			value: "restore",
		},
		DELETE: OperationLogOperationType{
			value: "delete",
		},
		SYNC: OperationLogOperationType{
			value: "sync",
		},
		VAULT_DELETE: OperationLogOperationType{
			value: "vault_delete",
		},
		REMOVE_RESOURCE: OperationLogOperationType{
			value: "remove_resource",
		},
	}
}

func (c OperationLogOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationLogOperationType) UnmarshalJSON(b []byte) error {
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

type OperationLogStatus struct {
	value string
}

type OperationLogStatusEnum struct {
	SUCCESS OperationLogStatus
	SKIPPED OperationLogStatus
	FAILED  OperationLogStatus
	RUNNING OperationLogStatus
	TIMEOUT OperationLogStatus
	WAITING OperationLogStatus
}

func GetOperationLogStatusEnum() OperationLogStatusEnum {
	return OperationLogStatusEnum{
		SUCCESS: OperationLogStatus{
			value: "success",
		},
		SKIPPED: OperationLogStatus{
			value: "skipped",
		},
		FAILED: OperationLogStatus{
			value: "failed",
		},
		RUNNING: OperationLogStatus{
			value: "running",
		},
		TIMEOUT: OperationLogStatus{
			value: "timeout",
		},
		WAITING: OperationLogStatus{
			value: "waiting",
		},
	}
}

func (c OperationLogStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationLogStatus) UnmarshalJSON(b []byte) error {
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
