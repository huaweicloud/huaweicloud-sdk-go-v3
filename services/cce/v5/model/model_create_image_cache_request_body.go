package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageCacheRequestBody 创建镜像缓存-respuest结构体。
type CreateImageCacheRequestBody struct {

	// **参数解释：** 镜像缓存名称。 **约束限制：** 不涉及 **取值范围：** 以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围1-128位，且不能以中划线(-)结尾。镜像缓存名称不可重复。 **默认取值：** 不涉及
	Name string `json:"name"`

	// 镜像缓存中的容器镜像列表。
	Images []string `json:"images"`

	// **参数解释：** 镜像缓存后端对应的存储盘大小，单位GiB。缓存对象为解压后的镜像文件，建议镜像缓存大小不低于该缓存中所有容器镜像大小总和的3倍。 **约束限制：** 不涉及 **取值范围：** [20-400] **默认取值：** 20
	ImageCacheSize *int32 `json:"image_cache_size,omitempty"`

	// **参数解释：** 镜像缓存有效的天数,不设置或值为0表示永久有效。镜像缓存到达有效期后自动过期并删除。 **约束限制：** 不涉及 **取值范围：** [0-10000] **默认取值：** 0
	RetentionDays *int32 `json:"retention_days,omitempty"`

	BuildingConfig *ImageCacheBuildingConfig `json:"building_config"`
}

func (o CreateImageCacheRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageCacheRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageCacheRequestBody", string(data)}, " ")
}
