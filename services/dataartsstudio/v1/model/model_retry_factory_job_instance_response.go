package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryFactoryJobInstanceResponse Response Object
type RetryFactoryJobInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryFactoryJobInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryFactoryJobInstanceResponse struct{}"
	}

	return strings.Join([]string{"RetryFactoryJobInstanceResponse", string(data)}, " ")
}
