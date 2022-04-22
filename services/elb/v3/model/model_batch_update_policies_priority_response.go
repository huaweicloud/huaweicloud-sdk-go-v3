package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchUpdatePoliciesPriorityResponse struct {

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdatePoliciesPriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityResponse", string(data)}, " ")
}
