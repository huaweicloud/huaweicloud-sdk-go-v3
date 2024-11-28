package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDesktopPoolActionRequest Request Object
type ExecuteDesktopPoolActionRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *DesktopPoolActionReq `json:"body,omitempty"`
}

func (o ExecuteDesktopPoolActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDesktopPoolActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDesktopPoolActionRequest", string(data)}, " ")
}
