package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReleaseParams struct {

	// 是否仅模拟安装过程
	DryRun *bool `json:"dry_run,omitempty"`

	// 是否允许重用已存在的名称
	Replace *bool `json:"replace,omitempty"`

	// 是否强制重新创建资源
	Recreate *bool `json:"recreate,omitempty"`

	// 是否禁止hook
	NoHooks *bool `json:"no_hooks,omitempty"`

	// 更新时重设values
	ResetValues *bool `json:"reset_values,omitempty"`

	// 发布资源的名称模板
	NameTemplate *string `json:"name_template,omitempty"`

	// 指定回滚版本号
	ReleaseVersion *int32 `json:"release_version,omitempty"`

	// 更新或删除时是否允许hook
	IncludeHooks *bool `json:"include_hooks,omitempty"`
}

func (o ReleaseParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseParams struct{}"
	}

	return strings.Join([]string{"ReleaseParams", string(data)}, " ")
}
