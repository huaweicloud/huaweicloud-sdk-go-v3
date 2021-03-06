package model

import (
	"encoding/json"

	"strings"
)

type CheckpointExtraInfoResp struct {
	// 备份名称

	Name *string `json:"name,omitempty"`
	// 备份描述

	Description *string `json:"description,omitempty"`
	// 备份保留天数

	RetentionDuration *int32 `json:"retention_duration,omitempty"`
}

func (o CheckpointExtraInfoResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckpointExtraInfoResp struct{}"
	}

	return strings.Join([]string{"CheckpointExtraInfoResp", string(data)}, " ")
}
