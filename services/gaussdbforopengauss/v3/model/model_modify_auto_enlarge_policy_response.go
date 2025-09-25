package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAutoEnlargePolicyResponse Response Object
type ModifyAutoEnlargePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"ModifyAutoEnlargePolicyResponse", string(data)}, " ")
}
