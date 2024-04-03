package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRemoteAssistanceRequest Request Object
type CancelRemoteAssistanceRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o CancelRemoteAssistanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRemoteAssistanceRequest struct{}"
	}

	return strings.Join([]string{"CancelRemoteAssistanceRequest", string(data)}, " ")
}
