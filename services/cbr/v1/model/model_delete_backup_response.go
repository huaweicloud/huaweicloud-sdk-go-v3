package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupResponse Response Object
type DeleteBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupResponse", string(data)}, " ")
}
