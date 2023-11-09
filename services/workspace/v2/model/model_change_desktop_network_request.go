package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDesktopNetworkRequest Request Object
type ChangeDesktopNetworkRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	Body *ChangeDesktopNetworkReq `json:"body,omitempty"`
}

func (o ChangeDesktopNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDesktopNetworkRequest struct{}"
	}

	return strings.Join([]string{"ChangeDesktopNetworkRequest", string(data)}, " ")
}
