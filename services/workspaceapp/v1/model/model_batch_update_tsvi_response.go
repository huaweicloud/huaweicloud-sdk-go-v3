package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTsviResponse Response Object
type BatchUpdateTsviResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateTsviResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTsviResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateTsviResponse", string(data)}, " ")
}
