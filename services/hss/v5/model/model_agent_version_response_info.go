package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentVersionResponseInfo agent版本信息详情列表
type AgentVersionResponseInfo struct {

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 最新版本的版本号
	LatestVersion *string `json:"latest_version,omitempty"`

	// 最新若干版本的版本号和版本说明
	VersionList *[]VersionList `json:"version_list,omitempty"`
}

func (o AgentVersionResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentVersionResponseInfo struct{}"
	}

	return strings.Join([]string{"AgentVersionResponseInfo", string(data)}, " ")
}
