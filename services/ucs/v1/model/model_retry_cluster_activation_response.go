package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryClusterActivationResponse Response Object
type RetryClusterActivationResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RetryClusterActivationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryClusterActivationResponse struct{}"
	}

	return strings.Join([]string{"RetryClusterActivationResponse", string(data)}, " ")
}
