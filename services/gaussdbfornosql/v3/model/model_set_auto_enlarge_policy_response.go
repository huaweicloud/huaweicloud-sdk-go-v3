package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoEnlargePolicyResponse Response Object
type SetAutoEnlargePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePolicyResponse", string(data)}, " ")
}
