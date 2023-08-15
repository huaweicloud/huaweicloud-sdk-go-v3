package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

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
	Schema         *string `json:"schema,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceUpdateImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberResponse", string(data)}, " ")
}
