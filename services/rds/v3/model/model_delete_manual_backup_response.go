package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteManualBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupResponse", string(data)}, " ")
}
