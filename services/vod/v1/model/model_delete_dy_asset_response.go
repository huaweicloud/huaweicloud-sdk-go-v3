package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDyAssetResponse Response Object
type DeleteDyAssetResponse struct {

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDyAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDyAssetResponse struct{}"
	}

	return strings.Join([]string{"DeleteDyAssetResponse", string(data)}, " ")
}
