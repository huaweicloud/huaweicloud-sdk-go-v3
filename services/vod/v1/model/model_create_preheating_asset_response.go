package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePreheatingAssetResponse Response Object
type CreatePreheatingAssetResponse struct {

	// 预热任务ID。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePreheatingAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePreheatingAssetResponse struct{}"
	}

	return strings.Join([]string{"CreatePreheatingAssetResponse", string(data)}, " ")
}
