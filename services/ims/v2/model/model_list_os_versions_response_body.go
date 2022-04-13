package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询操作系统列表响应体
type ListOsVersionsResponseBody struct {
	// 操作系统的平台值，如RedHat等

	Platform string `json:"platform"`
	// 操作系统的详情值

	VersionList []OsVersionInfo `json:"version_list"`
}

func (o ListOsVersionsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOsVersionsResponseBody struct{}"
	}

	return strings.Join([]string{"ListOsVersionsResponseBody", string(data)}, " ")
}
