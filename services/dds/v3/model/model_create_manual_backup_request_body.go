package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateManualBackupRequestBody struct {
	Backup *CreateManualBackupOption `json:"backup"`
}

func (o CreateManualBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateManualBackupRequestBody", string(data)}, " ")
}
