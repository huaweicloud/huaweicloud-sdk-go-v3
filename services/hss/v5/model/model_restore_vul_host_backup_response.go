package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreVulHostBackupResponse Response Object
type RestoreVulHostBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreVulHostBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreVulHostBackupResponse struct{}"
	}

	return strings.Join([]string{"RestoreVulHostBackupResponse", string(data)}, " ")
}
