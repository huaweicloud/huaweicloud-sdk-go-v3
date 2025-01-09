package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDesktopAttributesReq 查询云桌面列表响应
type ModifyDesktopAttributesReq struct {
	Desktop *ModifyDesktopAttributesReqDesktop `json:"desktop,omitempty"`
}

func (o ModifyDesktopAttributesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDesktopAttributesReq struct{}"
	}

	return strings.Join([]string{"ModifyDesktopAttributesReq", string(data)}, " ")
}
