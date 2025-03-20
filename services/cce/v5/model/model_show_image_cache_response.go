package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageCacheResponse Response Object
type ShowImageCacheResponse struct {
	ImageCache     *ImageCacheDetail `json:"image_cache,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowImageCacheResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageCacheResponse struct{}"
	}

	return strings.Join([]string{"ShowImageCacheResponse", string(data)}, " ")
}
