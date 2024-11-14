package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBackupResponse Response Object
type BatchDeleteBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBackupResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBackupResponse", string(data)}, " ")
}
