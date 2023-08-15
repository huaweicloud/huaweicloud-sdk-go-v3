package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupResponse Response Object
type ListBackupResponse struct {

	// 归档记录总数量
	Count *int64 `json:"count,omitempty"`

	// 归档记录列表
	Backups        *[]BackupDto `json:"backups,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupResponse struct{}"
	}

	return strings.Join([]string{"ListBackupResponse", string(data)}, " ")
}
