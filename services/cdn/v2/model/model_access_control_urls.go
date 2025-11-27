package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessControlUrls 需要解禁或封禁的URL信息
type AccessControlUrls struct {

	// URL必须带有“http://”或“https://”，单次最多输入1000个url。
	Urls []string `json:"urls"`

	// URL封禁天数，默认7天，取值范围1-30。
	BanDuration *int32 `json:"ban_duration,omitempty"`
}

func (o AccessControlUrls) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControlUrls struct{}"
	}

	return strings.Join([]string{"AccessControlUrls", string(data)}, " ")
}
