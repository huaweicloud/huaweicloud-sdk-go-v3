package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddProtectResponse struct {

	// 分支名称
	Name *string `json:"name,omitempty" xml:"name"`

	Commit *CommitRepoV2 `json:"commit,omitempty" xml:"commit"`

	// 是否保护
	Protected *bool `json:"protected,omitempty" xml:"protected"`

	// 是否允许开发者提交
	DevelopersCanPush *bool `json:"developers_can_push,omitempty" xml:"developers_can_push"`

	// 是否允许开发者合并
	DevelopersCanMerge *bool `json:"developers_can_merge,omitempty" xml:"developers_can_merge"`

	// 是否允许管理员提交
	MasterCanPush *bool `json:"master_can_push,omitempty" xml:"master_can_push"`

	// 是否允许管理员合并
	MasterCanMerge *bool `json:"master_can_merge,omitempty" xml:"master_can_merge"`

	// 没有人允许提交
	NoOneCanPush *bool `json:"no_one_can_push,omitempty" xml:"no_one_can_push"`

	// 没有人允许合并
	NoOneCanMerge *bool `json:"no_one_can_merge,omitempty" xml:"no_one_can_merge"`

	// 是否在一个打开的合并请求
	InAnOpenedMergeRequest *bool `json:"in_an_opened_merge_request,omitempty" xml:"in_an_opened_merge_request"`
}

func (o AddProtectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectResponse struct{}"
	}

	return strings.Join([]string{"AddProtectResponse", string(data)}, " ")
}
