package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteImageServerResponse Response Object
type BatchDeleteImageServerResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteImageServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteImageServerResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteImageServerResponse", string(data)}, " ")
}
