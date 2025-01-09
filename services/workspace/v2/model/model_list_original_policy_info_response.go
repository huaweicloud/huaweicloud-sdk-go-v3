package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOriginalPolicyInfoResponse Response Object
type ListOriginalPolicyInfoResponse struct {
	Policies       *Policies `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOriginalPolicyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOriginalPolicyInfoResponse struct{}"
	}

	return strings.Join([]string{"ListOriginalPolicyInfoResponse", string(data)}, " ")
}
