package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneImagesRequest Request Object
type ListCloudPhoneImagesRequest struct {

	// 镜像类型 公共镜像：public 私有镜像：private 共享镜像：share 所有类型镜像：all
	ImageType *string `json:"image_type,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的镜像个数。取值范围：1~500（默认值为500），一般设置为10、20、50。 当image_type传all时，分页返回顺序按公共镜像：public 私有镜像，private 共享镜像：share
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCloudPhoneImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneImagesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneImagesRequest", string(data)}, " ")
}
