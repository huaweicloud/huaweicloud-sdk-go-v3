package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopRequest Request Object
type UpdateDesktopRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	Body *ModifyDesktopAttributesReq `json:"body,omitempty"`
}

func (o UpdateDesktopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopRequest struct{}"
	}

	return strings.Join([]string{"UpdateDesktopRequest", string(data)}, " ")
}
