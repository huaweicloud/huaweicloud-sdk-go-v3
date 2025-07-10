package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopPoolRequest Request Object
type ResizeDesktopPoolRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *ResizeDesktopPoolReq `json:"body,omitempty"`
}

func (o ResizeDesktopPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopPoolRequest struct{}"
	}

	return strings.Join([]string{"ResizeDesktopPoolRequest", string(data)}, " ")
}
