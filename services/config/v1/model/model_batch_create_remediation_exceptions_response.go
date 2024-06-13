package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateRemediationExceptionsResponse Response Object
type BatchCreateRemediationExceptionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateRemediationExceptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateRemediationExceptionsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateRemediationExceptionsResponse", string(data)}, " ")
}
