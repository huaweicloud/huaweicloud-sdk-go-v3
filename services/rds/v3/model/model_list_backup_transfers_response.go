package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupTransfersResponse Response Object
type ListBackupTransfersResponse struct {

	// 任务数量
	Total *int64 `json:"total,omitempty"`

	// 任务列表
	TransferList   *[]BackupTransferInfo `json:"transfer_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListBackupTransfersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupTransfersResponse struct{}"
	}

	return strings.Join([]string{"ListBackupTransfersResponse", string(data)}, " ")
}
