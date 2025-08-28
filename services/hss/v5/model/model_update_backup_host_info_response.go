package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupHostInfoResponse Response Object
type UpdateBackupHostInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateBackupHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupHostInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateBackupHostInfoResponse", string(data)}, " ")
}
