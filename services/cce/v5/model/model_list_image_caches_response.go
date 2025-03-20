package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageCachesResponse Response Object
type ListImageCachesResponse struct {

	// 镜像缓存列表。
	ImageCaches    *[]ImageCacheDetail `json:"image_caches,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListImageCachesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageCachesResponse struct{}"
	}

	return strings.Join([]string{"ListImageCachesResponse", string(data)}, " ")
}
