package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachPolicyResponse Response Object
type DetachPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachPolicyResponse struct{}"
	}

	return strings.Join([]string{"DetachPolicyResponse", string(data)}, " ")
}
