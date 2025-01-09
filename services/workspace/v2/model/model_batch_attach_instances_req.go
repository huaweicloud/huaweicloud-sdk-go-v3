package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachInstancesReq 分配用户请求。
type BatchAttachInstancesReq struct {

	// 桌面信息列表。
	Desktops *[]AttachInstancesDesktopInfo `json:"desktops,omitempty"`

	// 用户信息列表。
	Users *[]AttachInstancesUserInfo `json:"users,omitempty"`

	AssignModel *AssignModelInfo `json:"assign_model,omitempty"`
}

func (o BatchAttachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInstancesReq struct{}"
	}

	return strings.Join([]string{"BatchAttachInstancesReq", string(data)}, " ")
}
