package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDownloadUrlResponse Response Object
type CreateDownloadUrlResponse struct {

	// OBS 预签名下载地址。
	DownloadUrl *string `json:"download_url,omitempty"`

	// OBS 桶名。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// OBS 对象键。
	ObsObjectKey *string `json:"obs_object_key,omitempty"`

	// 下载地址过期时间（ISO8601格式，UTC时区）。
	ExpiresAt *string `json:"expires_at,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDownloadUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDownloadUrlResponse struct{}"
	}

	return strings.Join([]string{"CreateDownloadUrlResponse", string(data)}, " ")
}
