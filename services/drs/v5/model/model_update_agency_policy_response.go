package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyPolicyResponse Response Object
type UpdateAgencyPolicyResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateAgencyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyPolicyResponse", string(data)}, " ")
}
