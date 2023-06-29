package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMemberReq 分享项目请求体
type UpdateMemberReq struct {
	Role *ProjectRoleType `json:"role"`
}

func (o UpdateMemberReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberReq struct{}"
	}

	return strings.Join([]string{"UpdateMemberReq", string(data)}, " ")
}
