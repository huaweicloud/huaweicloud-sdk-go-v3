package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopVolumesRequest Request Object
type AddDesktopVolumesRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	Body *AddVolumesReq `json:"body,omitempty"`
}

func (o AddDesktopVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopVolumesRequest struct{}"
	}

	return strings.Join([]string{"AddDesktopVolumesRequest", string(data)}, " ")
}
