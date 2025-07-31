package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceUpdateImageMemberResponse Response Object
type GlanceUpdateImageMemberResponse struct {

	// 共享状态
	Status *string `json:"status,omitempty"`

	// 共享时间，格式为UTC时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间，格式为UTC时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 成员ID
	MemberId *string `json:"member_id,omitempty"`

	// 共享视图
	Schema *string `json:"schema,omitempty"`

	// 共享成员类型
	MemberType *string `json:"member_type,omitempty"`

	// 共享组织的URN,仅当member_type类型为organization时，才会返回urn字段
	Urn            *string `json:"urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceUpdateImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberResponse", string(data)}, " ")
}
