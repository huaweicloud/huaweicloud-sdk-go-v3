package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopVolumesRequest Request Object
type DeleteDesktopVolumesRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	Body *DeleteVolumesReq `json:"body,omitempty"`
}

func (o DeleteDesktopVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopVolumesRequest struct{}"
	}

	return strings.Join([]string{"DeleteDesktopVolumesRequest", string(data)}, " ")
}
