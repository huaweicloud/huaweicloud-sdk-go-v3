package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryJobResponse Response Object
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
