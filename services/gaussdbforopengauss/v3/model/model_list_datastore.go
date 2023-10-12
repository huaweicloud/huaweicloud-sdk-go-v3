package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatastore 数据库信息。
type ListDatastore struct {

	// 数据库引擎。
	Type string `json:"type"`

	// 数据库大版本。
	Version string `json:"version"`

	// 数据库小版本。
	CompleteVersion *string `json:"complete_version,omitempty"`

	// 数据库已升级的热补丁版本，当数据库热补丁升级成功后，该值不为空。
	HotfixVersions *string `json:"hotfix_versions,omitempty"`
}

func (o ListDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatastore struct{}"
	}

	return strings.Join([]string{"ListDatastore", string(data)}, " ")
}
