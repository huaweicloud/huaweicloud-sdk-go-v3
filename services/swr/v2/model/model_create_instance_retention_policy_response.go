package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRetentionPolicyResponse Response Object
type CreateInstanceRetentionPolicyResponse struct {

	// 命名空间ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInstanceRetentionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRetentionPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceRetentionPolicyResponse", string(data)}, " ")
}
