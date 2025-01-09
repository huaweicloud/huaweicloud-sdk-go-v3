package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessAddressBackupConfigRequest Request Object
type ListAccessAddressBackupConfigRequest struct {
}

func (o ListAccessAddressBackupConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessAddressBackupConfigRequest struct{}"
	}

	return strings.Join([]string{"ListAccessAddressBackupConfigRequest", string(data)}, " ")
}
