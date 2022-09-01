package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackupsResponse struct {

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode"`

	// 备份总个数。请求失败时，字段为空。
	BackupCount *int32 `json:"backupCount,omitempty" xml:"backupCount"`

	// 当前Project ID下的所有图的备份列表。请求失败时，字段为空。
	BackupList     *[]Backup `json:"backupList,omitempty" xml:"backupList"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
