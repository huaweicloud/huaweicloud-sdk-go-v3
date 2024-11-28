package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRedisPitrPolicyResponse Response Object
type SetRedisPitrPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRedisPitrPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRedisPitrPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetRedisPitrPolicyResponse", string(data)}, " ")
}
