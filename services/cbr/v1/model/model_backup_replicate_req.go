package model

import (
	"encoding/json"

	"strings"
)

//
type BackupReplicateReq struct {
	Replicate *BackupReplicateReqBody `json:"replicate"`
}

func (o BackupReplicateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BackupReplicateReq struct{}"
	}

	return strings.Join([]string{"BackupReplicateReq", string(data)}, " ")
}
