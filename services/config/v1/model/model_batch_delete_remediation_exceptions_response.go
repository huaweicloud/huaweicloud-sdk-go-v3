package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRemediationExceptionsResponse Response Object
type BatchDeleteRemediationExceptionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteRemediationExceptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRemediationExceptionsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRemediationExceptionsResponse", string(data)}, " ")
}
