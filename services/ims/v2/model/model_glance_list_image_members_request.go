package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceListImageMembersRequest Request Object
type GlanceListImageMembersRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 查询镜像成员列表时每页的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识，用于查询下一页内容。
	Marker *string `json:"marker,omitempty"`
}

func (o GlanceListImageMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMembersRequest", string(data)}, " ")
}
