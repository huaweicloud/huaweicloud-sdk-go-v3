package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDbUserCommentRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 数据库用户名。
	UserName string `json:"user_name" xml:"user_name"`

	Body *UpdateDbUserReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateDbUserCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDbUserCommentRequest", string(data)}, " ")
}
