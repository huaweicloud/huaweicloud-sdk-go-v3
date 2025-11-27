package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyDefinationResponse Response Object
type UpdatePolicyDefinationResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdatePolicyDefinationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyDefinationResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyDefinationResponse", string(data)}, " ")
}
