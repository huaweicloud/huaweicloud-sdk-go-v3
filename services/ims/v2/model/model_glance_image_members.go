package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceImageMembers 获取镜像成员列表
type GlanceImageMembers struct {

	// 共享状态。
	Status string `json:"status"`

	// 共享时间，格式为UTC时间。
	CreatedAt string `json:"created_at"`

	// 更新时间，格式为UTC时间。
	UpdatedAt string `json:"updated_at"`

	// 镜像ID。
	ImageId string `json:"image_id"`

	// 成员ID。
	MemberId string `json:"member_id"`

	// 共享视图。
	Schema string `json:"schema"`

	// 共享成员类型。
	MemberType *string `json:"member_type,omitempty"`

	// 共享组织的URN仅当member_type类型为organization时，才会返回urn字段。
	Urn *string `json:"urn,omitempty"`
}

func (o GlanceImageMembers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceImageMembers struct{}"
	}

	return strings.Join([]string{"GlanceImageMembers", string(data)}, " ")
}
