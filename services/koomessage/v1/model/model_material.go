package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 模板中使用的素材。
type Material struct {

	// 素材ID。
	Id string `json:"id"`

	// 创建时间。
	CreatedAt string `json:"created_at"`

	// 资源类型。 - image：表示图片 - video：表示视频 - thumbnail：表示缩略图
	ResourceType string `json:"resource_type"`

	// 文件名称。
	FileName string `json:"file_name"`

	// 资源ID。
	AimResourceId string `json:"aim_resource_id"`

	// 从OBS返回的文件Key。
	ObsObjectKey string `json:"obs_object_key"`

	// 文件访问路径。
	ObsFileUrl *string `json:"obs_file_url,omitempty"`

	// 图像比例。 - oneToOne：指1:1比例 - sixteenToNine：指16:9比例 - threeToOne：指3:1比例 - fortyEightToSixtyFive：指48:65比例 - twentyOneToNine：指21:9比例
	ImageRate *string `json:"image_rate,omitempty"`

	// 素材详细描述。
	Description *string `json:"description,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
}

func (o Material) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Material struct{}"
	}

	return strings.Join([]string{"Material", string(data)}, " ")
}
