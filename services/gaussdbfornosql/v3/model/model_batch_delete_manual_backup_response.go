package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteManualBackupResponse Response Object
type BatchDeleteManualBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteManualBackupResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteManualBackupResponse", string(data)}, " ")
}
