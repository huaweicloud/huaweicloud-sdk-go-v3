package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferBackupRequest Request Object
type TransferBackupRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *TransferBackupRequestBody `json:"body,omitempty"`
}

func (o TransferBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferBackupRequest struct{}"
	}

	return strings.Join([]string{"TransferBackupRequest", string(data)}, " ")
}
