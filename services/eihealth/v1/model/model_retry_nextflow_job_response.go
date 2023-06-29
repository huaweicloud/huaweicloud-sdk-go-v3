package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryNextflowJobResponse Response Object
type RetryNextflowJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryNextflowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryNextflowJobResponse struct{}"
	}

	return strings.Join([]string{"RetryNextflowJobResponse", string(data)}, " ")
}
