package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScreenRecordsResponse Response Object
type BatchDeleteScreenRecordsResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteScreenRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScreenRecordsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteScreenRecordsResponse", string(data)}, " ")
}
