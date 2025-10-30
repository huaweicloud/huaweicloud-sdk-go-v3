package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryTaskResponse Response Object
type RetryTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskResponse struct{}"
	}

	return strings.Join([]string{"RetryTaskResponse", string(data)}, " ")
}
