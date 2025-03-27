package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCopyStateResponse Response Object
type UpdateCopyStateResponse struct {

	// 更新任务对应源端复制状态成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCopyStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCopyStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCopyStateResponse", string(data)}, " ")
}
