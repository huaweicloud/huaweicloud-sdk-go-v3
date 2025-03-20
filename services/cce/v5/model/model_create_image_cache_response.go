package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageCacheResponse Response Object
type CreateImageCacheResponse struct {
	ImageCache     *ImageCacheDetail `json:"image_cache,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateImageCacheResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageCacheResponse struct{}"
	}

	return strings.Join([]string{"CreateImageCacheResponse", string(data)}, " ")
}
