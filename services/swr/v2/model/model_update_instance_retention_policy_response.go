package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRetentionPolicyResponse Response Object
type UpdateInstanceRetentionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceRetentionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRetentionPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRetentionPolicyResponse", string(data)}, " ")
}
