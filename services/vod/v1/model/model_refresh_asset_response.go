package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshAssetResponse Response Object
type RefreshAssetResponse struct {

	// 刷新任务ID。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RefreshAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshAssetResponse struct{}"
	}

	return strings.Join([]string{"RefreshAssetResponse", string(data)}, " ")
}
