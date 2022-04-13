package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGraphBackupsResponse struct {
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。

	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。

	ErrorCode *string `json:"errorCode,omitempty"`
	// 备份总个数。请求失败时。字段为空。

	BackupCount *int32 `json:"backupCount,omitempty"`
	// 当前Project下指定Graph的的备份列表。请求失败时，字段为空。

	BackupList     *[]Backup `json:"backupList,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGraphBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListGraphBackupsResponse", string(data)}, " ")
}
