package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoEnlargePoliciesResponse Response Object
type SetAutoEnlargePoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAutoEnlargePoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePoliciesResponse struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePoliciesResponse", string(data)}, " ")
}
