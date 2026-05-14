package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetEditTaskResponse Response Object
type CreateAssetEditTaskResponse struct {

	// 编辑任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 媒资ID
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetEditTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetEditTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetEditTaskResponse", string(data)}, " ")
}
