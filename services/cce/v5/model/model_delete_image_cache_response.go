package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageCacheResponse Response Object
type DeleteImageCacheResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteImageCacheResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageCacheResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageCacheResponse", string(data)}, " ")
}
