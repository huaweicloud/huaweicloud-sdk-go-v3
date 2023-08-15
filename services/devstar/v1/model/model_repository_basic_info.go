package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryBasicInfo struct {

	// 仓库uuid
	Uuid *string `json:"uuid,omitempty"`

	// 仓库id
	Id *string `json:"id,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 仓库git的https下载地址
	HttpsUrl *string `json:"https_url,omitempty"`

	// 仓库git的ssh下载地址
	SshUrl *string `json:"ssh_url,omitempty"`

	// 仓库codehub内容浏览页面地址
	WebUrl *string `json:"web_url,omitempty"`
}

func (o RepositoryBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryBasicInfo struct{}"
	}

	return strings.Join([]string{"RepositoryBasicInfo", string(data)}, " ")
}
