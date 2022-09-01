package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GlanceUpdateImageMemberResponse struct {

	// 共享状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 共享时间，格式为UTC时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 更新时间，格式为UTC时间
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`

	// 成员ID
	MemberId *string `json:"member_id,omitempty" xml:"member_id"`

	// 共享视图
	Schema         *string `json:"schema,omitempty" xml:"schema"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceUpdateImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberResponse", string(data)}, " ")
}
