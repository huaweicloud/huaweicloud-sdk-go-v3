package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopPoolRequest Request Object
type UpdateDesktopPoolRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *UpdateDesktopPoolAttributesReq `json:"body,omitempty"`
}

func (o UpdateDesktopPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopPoolRequest struct{}"
	}

	return strings.Join([]string{"UpdateDesktopPoolRequest", string(data)}, " ")
}
