package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InitExecuteTaskInfo struct {

	// 版本信息
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 分支/迭代uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 是否只查询，不初始化（如果不存在）
	IsQuery *bool `json:"is_query,omitempty"`
}

func (o InitExecuteTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitExecuteTaskInfo struct{}"
	}

	return strings.Join([]string{"InitExecuteTaskInfo", string(data)}, " ")
}
