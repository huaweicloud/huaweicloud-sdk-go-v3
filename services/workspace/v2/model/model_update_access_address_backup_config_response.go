package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessAddressBackupConfigResponse Response Object
type UpdateAccessAddressBackupConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAccessAddressBackupConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessAddressBackupConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessAddressBackupConfigResponse", string(data)}, " ")
}
