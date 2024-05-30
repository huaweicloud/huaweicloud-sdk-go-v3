package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneImagesRequest Request Object
type ListCloudPhoneImagesRequest struct {

	// 镜像类型 公共镜像：public 私有镜像：private 共享镜像：share 所有类型镜像：all
	ImageType *string `json:"image_type,omitempty"`
}

func (o ListCloudPhoneImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneImagesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneImagesRequest", string(data)}, " ")
}
