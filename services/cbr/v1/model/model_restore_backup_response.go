package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestoreBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBackupResponse struct{}"
	}

	return strings.Join([]string{"RestoreBackupResponse", string(data)}, " ")
}
