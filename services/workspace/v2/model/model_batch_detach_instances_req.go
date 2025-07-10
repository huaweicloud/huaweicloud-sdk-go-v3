package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachInstancesReq 批量解绑用户请求。
type BatchDetachInstancesReq struct {

	// 解绑的桌面信息列表。
	Desktops *[]DetachInstancesDesktopInfo `json:"desktops,omitempty"`
}

func (o BatchDetachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInstancesReq struct{}"
	}

	return strings.Join([]string{"BatchDetachInstancesReq", string(data)}, " ")
}
