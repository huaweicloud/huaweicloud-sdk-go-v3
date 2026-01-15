package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInstancesReq 解绑用户请求。
type DetachInstancesReq struct {

	// 桌面id列表，不能为空。运行状态为【已休眠】的桌面不支持解绑操作。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 解绑后是否关机。
	ShutoffAfterDetach *bool `json:"shutoff_after_detach,omitempty"`
}

func (o DetachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInstancesReq struct{}"
	}

	return strings.Join([]string{"DetachInstancesReq", string(data)}, " ")
}
