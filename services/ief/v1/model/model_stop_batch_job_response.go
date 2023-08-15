package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBatchJobResponse Response Object
type StopBatchJobResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBatchJobResponse struct{}"
	}

	return strings.Join([]string{"StopBatchJobResponse", string(data)}, " ")
}
