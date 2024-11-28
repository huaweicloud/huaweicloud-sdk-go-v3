package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLivePlatformsResponse Response Object
type ListLivePlatformsResponse struct {

	// 直播平台列表
	LivePlatforms *[]LivePlatformInfo `json:"live_platforms,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLivePlatformsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLivePlatformsResponse struct{}"
	}

	return strings.Join([]string{"ListLivePlatformsResponse", string(data)}, " ")
}
