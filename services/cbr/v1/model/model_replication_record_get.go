package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type ReplicationRecordGet struct {
	// 复制的开始时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 复制的目的备份ID

	DestinationBackupId *string `json:"destination_backup_id,omitempty"`
	// 复制的目的备份记录ID

	DestinationCheckpointId *string `json:"destination_checkpoint_id,omitempty"`
	// 复制的目标项目ID

	DestinationProjectId *string `json:"destination_project_id,omitempty"`
	// 复制的目标区域

	DestinationRegion *string `json:"destination_region,omitempty"`
	// 目标存储库ID

	DestinationVaultId *string `json:"destination_vault_id,omitempty"`
	// 复制附加信息

	ExtraInfo *string `json:"extra_info,omitempty"`
	// 复制记录ID

	Id string `json:"id"`
	// 复制的源备份ID

	SourceBackupId *string `json:"source_backup_id,omitempty"`
	// 复制的源备份记录ID

	SourceCheckpointId *string `json:"source_checkpoint_id,omitempty"`
	// 复制的源项目ID

	SourceProjectId *string `json:"source_project_id,omitempty"`
	// 复制的源区域

	SourceRegion *string `json:"source_region,omitempty"`
	// 复制的状态

	Status *ReplicationRecordGetStatus `json:"status,omitempty"`
	// 备份所在的存储库ID

	VaultId *string `json:"vault_id,omitempty"`
}

func (o ReplicationRecordGet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationRecordGet struct{}"
	}

	return strings.Join([]string{"ReplicationRecordGet", string(data)}, " ")
}

type ReplicationRecordGetStatus struct {
	value string
}

type ReplicationRecordGetStatusEnum struct {
	REPLICATING       ReplicationRecordGetStatus
	SUCCESS           ReplicationRecordGetStatus
	FAIL              ReplicationRecordGetStatus
	SKIP              ReplicationRecordGetStatus
	WAITING_REPLICATE ReplicationRecordGetStatus
}

func GetReplicationRecordGetStatusEnum() ReplicationRecordGetStatusEnum {
	return ReplicationRecordGetStatusEnum{
		REPLICATING: ReplicationRecordGetStatus{
			value: "replicating",
		},
		SUCCESS: ReplicationRecordGetStatus{
			value: "success",
		},
		FAIL: ReplicationRecordGetStatus{
			value: "fail",
		},
		SKIP: ReplicationRecordGetStatus{
			value: "skip",
		},
		WAITING_REPLICATE: ReplicationRecordGetStatus{
			value: "waiting_replicate",
		},
	}
}

func (c ReplicationRecordGetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReplicationRecordGetStatus) UnmarshalJSON(b []byte) error {
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
