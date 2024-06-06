package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRecyclePolicyResponse Response Object
type SetRecyclePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyResponse", string(data)}, " ")
}
