package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageMembersView 列举镜像成员响应体
type ListImageMembersView struct {

	// 共享时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 被共享账号的PROJECT_ID
	MemberId *string `json:"member_id,omitempty"`
}

func (o ListImageMembersView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageMembersView struct{}"
	}

	return strings.Join([]string{"ListImageMembersView", string(data)}, " ")
}
