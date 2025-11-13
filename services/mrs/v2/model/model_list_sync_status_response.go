package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncStatusResponse Response Object
type ListSyncStatusResponse struct {

	// 表示IAM同步任务是否正在进行中
	IsSynchronizing *bool `json:"is_synchronizing,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListSyncStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncStatusResponse struct{}"
	}

	return strings.Join([]string{"ListSyncStatusResponse", string(data)}, " ")
}
