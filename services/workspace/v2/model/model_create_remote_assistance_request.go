package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRemoteAssistanceRequest Request Object
type CreateRemoteAssistanceRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o CreateRemoteAssistanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemoteAssistanceRequest struct{}"
	}

	return strings.Join([]string{"CreateRemoteAssistanceRequest", string(data)}, " ")
}
