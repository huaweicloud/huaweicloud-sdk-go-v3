package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDesktopDetailRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o ShowDesktopDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopDetailRequest", string(data)}, " ")
}
