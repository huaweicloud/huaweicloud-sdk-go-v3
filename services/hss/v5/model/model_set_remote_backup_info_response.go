package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRemoteBackupInfoResponse Response Object
type SetRemoteBackupInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRemoteBackupInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRemoteBackupInfoResponse struct{}"
	}

	return strings.Join([]string{"SetRemoteBackupInfoResponse", string(data)}, " ")
}
