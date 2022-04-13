package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceShowImageMemberResponse struct {
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

func (o GlanceShowImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberResponse", string(data)}, " ")
}
