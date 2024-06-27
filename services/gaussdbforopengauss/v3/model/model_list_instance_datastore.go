package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDatastore 数据库信息。
type ListInstanceDatastore struct {

	// 数据库引擎。
	Type string `json:"type"`

	// 数据库大版本。
	Version string `json:"version"`

	// 数据库小版本。
	CompleteVersion *string `json:"complete_version,omitempty"`

	// 数据库正在升级的目标版本。
	TargetVersion *string `json:"target_version,omitempty"`

	// 数据库内核版本
	CompleteKernelVersion *string `json:"complete_kernel_version,omitempty"`

	// 热补丁信息列表
	HotfixVersionInfos *[]HotfixVersionInfo `json:"hotfix_version_infos,omitempty"`
}

func (o ListInstanceDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDatastore struct{}"
	}

	return strings.Join([]string{"ListInstanceDatastore", string(data)}, " ")
}
