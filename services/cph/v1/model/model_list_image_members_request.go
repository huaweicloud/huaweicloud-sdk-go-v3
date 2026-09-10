package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageMembersRequest Request Object
type ListImageMembersRequest struct {

	// 镜像id。
	ImageId string `json:"image_id"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的共享账号个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListImageMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageMembersRequest struct{}"
	}

	return strings.Join([]string{"ListImageMembersRequest", string(data)}, " ")
}
