package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePropagationRoutePolicyResponse Response Object
type UpdatePropagationRoutePolicyResponse struct {
	Propagation *Propagation `json:"propagation,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePropagationRoutePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePropagationRoutePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePropagationRoutePolicyResponse", string(data)}, " ")
}
