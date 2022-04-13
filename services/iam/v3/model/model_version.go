package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Version struct {
	// 版本状态。

	Status string `json:"status"`
	// 最后更新时间。

	Updated string `json:"updated"`
	// 版本的资源链接信息。

	Links []VersionLinks `json:"links"`
	// 版本号，如v3.6。

	Id string `json:"id"`
	// 支持的消息格式。

	MediaTypes []VersionMediatypes `json:"media-types"`
}

func (o Version) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Version struct{}"
	}

	return strings.Join([]string{"Version", string(data)}, " ")
}
