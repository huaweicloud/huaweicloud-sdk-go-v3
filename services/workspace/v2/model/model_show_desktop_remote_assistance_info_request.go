package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopRemoteAssistanceInfoRequest Request Object
type ShowDesktopRemoteAssistanceInfoRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o ShowDesktopRemoteAssistanceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopRemoteAssistanceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopRemoteAssistanceInfoRequest", string(data)}, " ")
}
