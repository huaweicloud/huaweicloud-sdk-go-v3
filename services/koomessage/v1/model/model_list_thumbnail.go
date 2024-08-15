package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThumbnail 视频封面图列表。
type ListThumbnail struct {

	// 视频封面图ID。
	Id string `json:"id"`

	// 创建时间。
	CreatedAt string `json:"created_at"`

	// 是否作为视频素材封面。 - 0：否 - 1：是
	IsPrimary int32 `json:"is_primary"`

	// 文件名称。
	FileName string `json:"file_name"`

	// 资源ID。
	AimResourceId string `json:"aim_resource_id"`

	// 从OBS返回的文件Key。
	ObsObjectKey string `json:"obs_object_key"`

	// 图像比例。 - oneToOne：指1:1比例 - sixteenToNine：指16:9比例 - threeToOne：指3:1比例 - fortyEightToSixtyFive：指48:65比例 - twentyOneToNine：指21:9比例 - threeToFour：指3:4比例
	ImageRate *string `json:"image_rate,omitempty"`

	// 视频封面图是否自动从系统生成。 - 0：系统自动生成 - 1：上传自定义
	IsAutoGen *int32 `json:"is_auto_gen,omitempty"`

	// 视频封面图的详细描述。
	Description *string `json:"description,omitempty"`

	// OBS桶名称。
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 素材所占空间大小。
	Size *int32 `json:"size,omitempty"`

	// 文件访问路径。
	ObsFileUrl *string `json:"obs_file_url,omitempty"`
}

func (o ListThumbnail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThumbnail struct{}"
	}

	return strings.Join([]string{"ListThumbnail", string(data)}, " ")
}
