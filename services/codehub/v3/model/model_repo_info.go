package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoInfo struct {

	// 创建时间
	CreateAt *string `json:"createAt,omitempty" xml:"createAt"`

	// 仓库组名
	GroupName *string `json:"groupName,omitempty" xml:"groupName"`

	// https url
	HttpUrl *string `json:"httpUrl,omitempty" xml:"httpUrl"`

	// 仓库uuid
	Id *string `json:"id,omitempty" xml:"id"`

	// 仓库名
	Name *string `json:"name,omitempty" xml:"name"`

	// 项目的uuid
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`

	// 项目是否被删除
	ProjectIsDelete *string `json:"projectIsDelete,omitempty" xml:"projectIsDelete"`

	// 仓库主键id
	RepoId *string `json:"repoId,omitempty" xml:"repoId"`

	// ssh url
	SshUrl *string `json:"sshUrl,omitempty" xml:"sshUrl"`

	// 是否可见：0私有仓库，20公有仓库
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel"`

	// web url 访问路径
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl"`
}

func (o RepoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoInfo struct{}"
	}

	return strings.Join([]string{"RepoInfo", string(data)}, " ")
}
