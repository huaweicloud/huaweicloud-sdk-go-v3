package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataPolicyResponse Response Object
type UpdateDataPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDataPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataPolicyResponse", string(data)}, " ")
}
