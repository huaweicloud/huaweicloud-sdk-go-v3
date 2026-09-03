package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDesktopPoolUserRequest Request Object
type AttachDesktopPoolUserRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *AttachDesktopPoolUserReq `json:"body,omitempty"`
}

func (o AttachDesktopPoolUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDesktopPoolUserRequest struct{}"
	}

	return strings.Join([]string{"AttachDesktopPoolUserRequest", string(data)}, " ")
}
