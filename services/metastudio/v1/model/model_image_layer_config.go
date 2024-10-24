package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageLayerConfig 素材图片图层配置。
type ImageLayerConfig struct {

	// **参数解释**： 图片文件的URL。 **约束限制**： * 仅直播支持外部URL，其他业务通过资产库查询获取，不支持外部URL。 **取值范围**： 字符长度1-2048位。 **默认取值**： 不涉及
	ImageUrl *string `json:"image_url,omitempty"`
}

func (o ImageLayerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageLayerConfig struct{}"
	}

	return strings.Join([]string{"ImageLayerConfig", string(data)}, " ")
}
