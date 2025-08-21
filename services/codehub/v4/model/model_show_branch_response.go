package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBranchResponse Response Object
type ShowBranchResponse struct {

	// 分支名
	Name *string `json:"name,omitempty"`

	Commit *CommitDto `json:"commit,omitempty"`

	// 用户id
	Merged *bool `json:"merged,omitempty"`

	// 是否为保护分支
	Protected *bool `json:"protected,omitempty"`

	// 开发者是否可以推送
	DevelopersCanPush *bool `json:"developers_can_push,omitempty"`

	// 开发者是否可以合并
	DevelopersCanMerge *bool `json:"developers_can_merge,omitempty"`

	// 是否可以推送
	CanPush *bool `json:"can_push,omitempty"`

	// 是否为默认分支
	Default *bool `json:"default,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBranchResponse struct{}"
	}

	return strings.Join([]string{"ShowBranchResponse", string(data)}, " ")
}
