package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreBatchAttachInstancesReq 分配用户请求。
type PreBatchAttachInstancesReq struct {

	// 桌面信息列表。
	Desktops *[]AttachInstancesDesktopInfo `json:"desktops,omitempty"`

	// 用户信息列表。
	Users *[]AttachInstancesUserInfo `json:"users,omitempty"`

	AssignModel *AssignModelInfo `json:"assign_model,omitempty"`
}

func (o PreBatchAttachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreBatchAttachInstancesReq struct{}"
	}

	return strings.Join([]string{"PreBatchAttachInstancesReq", string(data)}, " ")
}
