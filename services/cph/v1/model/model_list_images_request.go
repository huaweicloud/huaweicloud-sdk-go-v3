package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesRequest Request Object
type ListImagesRequest struct {

	// 镜像类型。 - public：公共镜像 - private：私有镜像 - share：共享镜像
	ImageType *string `json:"image_type,omitempty"`

	// 镜像状态。 - 0：CREATING 创建中 - 1：PRODUCTION 生产态，可使用 - 2：CREATE_FAILED 创建失败
	Status *int32 `json:"status,omitempty"`

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`

	// 镜像id
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 起始时间
	CreateSince *int64 `json:"create_since,omitempty"`

	// 截止时间
	CreateUntil *int64 `json:"create_until,omitempty"`

	// 共享镜像账号的projectId
	SrcProjectId *string `json:"src_project_id,omitempty"`
}

func (o ListImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesRequest struct{}"
	}

	return strings.Join([]string{"ListImagesRequest", string(data)}, " ")
}
