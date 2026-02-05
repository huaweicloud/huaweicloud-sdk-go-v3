package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupSelectionRequestBody 是否删除或保留自动备份请求体。
type DeleteBackupSelectionRequestBody struct {

	// 选择是否保留自动备份标志
	Selection bool `json:"selection"`
}

func (o DeleteBackupSelectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupSelectionRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteBackupSelectionRequestBody", string(data)}, " ")
}
