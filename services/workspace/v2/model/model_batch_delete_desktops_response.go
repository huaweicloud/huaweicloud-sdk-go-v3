package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopsResponse Response Object
type BatchDeleteDesktopsResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopsResponse", string(data)}, " ")
}
