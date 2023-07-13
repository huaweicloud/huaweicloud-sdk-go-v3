package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberRequest Request Object
type UpdateMemberRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 更新或者添加项目成员角色的用户id
	UserId string `json:"user_id"`

	Body *UpdateMemberReq `json:"body,omitempty"`
}

func (o UpdateMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequest", string(data)}, " ")
}
