package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessAddressBackupConfigRequest Request Object
type UpdateAccessAddressBackupConfigRequest struct {
	Body *UpdateAccessAddressBackupConfigReq `json:"body,omitempty"`
}

func (o UpdateAccessAddressBackupConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessAddressBackupConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccessAddressBackupConfigRequest", string(data)}, " ")
}
