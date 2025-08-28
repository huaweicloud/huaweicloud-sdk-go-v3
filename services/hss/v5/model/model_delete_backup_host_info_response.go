package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupHostInfoResponse Response Object
type DeleteBackupHostInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupHostInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupHostInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupHostInfoResponse", string(data)}, " ")
}
