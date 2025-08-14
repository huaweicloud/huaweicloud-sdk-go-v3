package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePreBootPolicyResponse Response Object
type UpdatePreBootPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePreBootPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePreBootPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePreBootPolicyResponse", string(data)}, " ")
}
