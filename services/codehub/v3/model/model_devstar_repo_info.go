package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevstarRepoInfo struct {

	// 仓库的uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid"`

	// 仓库ID
	RepoId *string `json:"repo_id,omitempty" xml:"repo_id"`

	// 仓库名称
	RepoName *string `json:"repo_name,omitempty" xml:"repo_name"`

	// 仓库SSH地址
	SshUrl *string `json:"ssh_url,omitempty" xml:"ssh_url"`

	// 仓库HTTPS地址
	CodeUrl *string `json:"code_url,omitempty" xml:"code_url"`

	// 仓库预览地址
	DetailUrl *string `json:"detail_url,omitempty" xml:"detail_url"`
}

func (o DevstarRepoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevstarRepoInfo struct{}"
	}

	return strings.Join([]string{"DevstarRepoInfo", string(data)}, " ")
}
