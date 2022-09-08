package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoInfo struct {

	// 创建时间
	CreateAt *string `json:"createAt,omitempty"`

	// 仓库组名(克隆地址中域名后面项目名前的一段 示例：git@codehub.alpha.devcloud.inhuawei.com:Demo00228/testword.git  组名：Demo00228)
	GroupName *string `json:"groupName,omitempty"`

	// https url
	HttpUrl *string `json:"httpUrl,omitempty"`

	// 仓库uuid(由CreateRepository接口返回)
	Id *string `json:"id,omitempty"`

	// 仓库名
	Name *string `json:"name,omitempty"`

	// 项目的uuid
	ProjectId *string `json:"projectId,omitempty"`

	// 项目是否被删除
	ProjectIsDelete *string `json:"projectIsDelete,omitempty"`

	// 仓库主键id
	RepoId *string `json:"repoId,omitempty"`

	// ssh url
	SshUrl *string `json:"sshUrl,omitempty"`

	// 是否可见：0私有仓库，20公有仓库
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty"`

	// web url 访问路径
	WebUrl *string `json:"webUrl,omitempty"`
}

func (o RepoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoInfo struct{}"
	}

	return strings.Join([]string{"RepoInfo", string(data)}, " ")
}
