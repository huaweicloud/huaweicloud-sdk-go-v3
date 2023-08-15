package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryBatchJobResponse Response Object
type RetryBatchJobResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RetryBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryBatchJobResponse struct{}"
	}

	return strings.Join([]string{"RetryBatchJobResponse", string(data)}, " ")
}
