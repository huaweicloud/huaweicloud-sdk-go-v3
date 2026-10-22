package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPresignedUrlResponse Response Object
type ShowPresignedUrlResponse struct {

	// 预签名URL
	PresignedUrl *string `json:"presigned_url,omitempty"`

	// 过期时间（秒）
	Expire         *int64 `json:"expire,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPresignedUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPresignedUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowPresignedUrlResponse", string(data)}, " ")
}
