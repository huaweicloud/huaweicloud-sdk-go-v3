package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSysprepInfoRequest Request Object
type ShowSysprepInfoRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o ShowSysprepInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSysprepInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSysprepInfoRequest", string(data)}, " ")
}
