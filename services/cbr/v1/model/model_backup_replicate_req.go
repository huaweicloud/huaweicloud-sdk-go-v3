package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupReplicateReq
type BackupReplicateReq struct {
	Replicate *BackupReplicateReqBody `json:"replicate"`
}

func (o BackupReplicateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupReplicateReq struct{}"
	}

	return strings.Join([]string{"BackupReplicateReq", string(data)}, " ")
}
