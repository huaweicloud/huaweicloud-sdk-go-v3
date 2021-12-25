package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetMultiTaskOffsetResponse struct {
	// 任务重置结果

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetMultiTaskOffsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMultiTaskOffsetResponse struct{}"
	}

	return strings.Join([]string{"ResetMultiTaskOffsetResponse", string(data)}, " ")
}
