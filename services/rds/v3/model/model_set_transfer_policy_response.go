package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTransferPolicyResponse Response Object
type SetTransferPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetTransferPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTransferPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetTransferPolicyResponse", string(data)}, " ")
}
