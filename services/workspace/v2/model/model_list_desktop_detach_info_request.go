package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopDetachInfoRequest Request Object
type ListDesktopDetachInfoRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o ListDesktopDetachInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopDetachInfoRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopDetachInfoRequest", string(data)}, " ")
}
