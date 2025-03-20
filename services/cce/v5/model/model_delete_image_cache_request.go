package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageCacheRequest Request Object
type DeleteImageCacheRequest struct {

	// **参数解释：** 镜像缓存ID。 **约束限制：** 不涉及 **取值范围：** 镜像缓存ID。 **默认取值：** 不涉及
	ImageCacheId string `json:"image_cache_id"`
}

func (o DeleteImageCacheRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageCacheRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageCacheRequest", string(data)}, " ")
}
