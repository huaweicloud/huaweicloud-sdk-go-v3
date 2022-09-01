package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRepositoryArchiveRequest struct {

	// 仓库的uuid
	RepositoryUuid string `json:"repository_uuid" xml:"repository_uuid"`

	// 分支名称
	Sha string `json:"sha" xml:"sha"`

	// 下载的压缩包格式
	Format string `json:"format" xml:"format"`
}

func (o ShowRepositoryArchiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryArchiveRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryArchiveRequest", string(data)}, " ")
}
