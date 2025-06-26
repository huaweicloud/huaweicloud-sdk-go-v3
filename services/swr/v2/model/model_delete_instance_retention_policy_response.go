package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceRetentionPolicyResponse Response Object
type DeleteInstanceRetentionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceRetentionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceRetentionPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRetentionPolicyResponse", string(data)}, " ")
}
