package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInstancesReq 解绑用户请求
type DetachInstancesReq struct {

	// 桌面id
	DesktopIds *[]string `json:"desktop_ids,omitempty"`
}

func (o DetachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInstancesReq struct{}"
	}

	return strings.Join([]string{"DetachInstancesReq", string(data)}, " ")
}
