package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RetryJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RetryJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryJobResponse struct{}"
	}

	return strings.Join([]string{"RetryJobResponse", string(data)}, " ")
}
