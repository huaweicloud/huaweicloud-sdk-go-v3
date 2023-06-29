package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataPathPolicyResponse Response Object
type UpdateDataPathPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDataPathPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataPathPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataPathPolicyResponse", string(data)}, " ")
}
