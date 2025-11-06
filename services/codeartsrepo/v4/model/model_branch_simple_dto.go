package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BranchSimpleDto struct {

	// 分支名称
	Name *string `json:"name,omitempty"`

	// 是否为默认分支
	Default *bool `json:"default,omitempty"`

	// 用户是否为默认分支
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否为默认分支
	CanRead *bool `json:"can_read,omitempty"`

	// 是否为默认分支
	CanDownload *bool `json:"can_download,omitempty"`

	// 是否为默认分支
	CanPush *bool `json:"can_push,omitempty"`
}

func (o BranchSimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BranchSimpleDto struct{}"
	}

	return strings.Join([]string{"BranchSimpleDto", string(data)}, " ")
}
