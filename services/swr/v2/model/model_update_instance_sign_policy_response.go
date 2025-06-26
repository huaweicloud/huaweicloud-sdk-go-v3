package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceSignPolicyResponse Response Object
type UpdateInstanceSignPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceSignPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceSignPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceSignPolicyResponse", string(data)}, " ")
}
