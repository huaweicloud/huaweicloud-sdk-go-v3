package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageMembersRequest Request Object
type ListImageMembersRequest struct {

	// 镜像id
	ImageId string `json:"image_id"`

	// 查询镜像成员列表时每页的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识，用于查询下一页内容。
	Marker *string `json:"marker,omitempty"`
}

func (o ListImageMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"ListImageMembersRequest", string(data)}, " ")
}
