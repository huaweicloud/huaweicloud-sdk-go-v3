package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPasswordRequest Request Object
type ResetPasswordRequest struct {

	// 指定待重置密码集群的ID
	ClusterId string `json:"cluster_id"`

	Body *ResetPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetPasswordRequest", string(data)}, " ")
}
