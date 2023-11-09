package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagByDesktopIdRequest Request Object
type ShowTagByDesktopIdRequest struct {

	// 桌面id。
	DesktopId string `json:"desktop_id"`
}

func (o ShowTagByDesktopIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagByDesktopIdRequest struct{}"
	}

	return strings.Join([]string{"ShowTagByDesktopIdRequest", string(data)}, " ")
}
