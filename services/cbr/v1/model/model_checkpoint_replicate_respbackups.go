package model

import (
	"encoding/json"

	"strings"
)

type CheckpointReplicateRespbackups struct {
	// 待复制的备份ID

	BackupId string `json:"backup_id"`
	// 复制记录ID

	ReplicationRecordId string `json:"replication_record_id"`
}

func (o CheckpointReplicateRespbackups) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckpointReplicateRespbackups struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateRespbackups", string(data)}, " ")
}
