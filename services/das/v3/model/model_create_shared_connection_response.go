package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSharedConnectionResponse Response Object
type CreateSharedConnectionResponse struct {

	// 状态。取值范围：true（成功）、false（失败）
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateSharedConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateSharedConnectionResponse", string(data)}, " ")
}
