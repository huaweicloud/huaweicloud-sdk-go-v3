package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadArchiveRequest Request Object
type DownloadArchiveRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支名、tag名、提交ID。
	Sha *string `json:"sha,omitempty"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过100000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 下载格式。 **取值范围：** - zip。 - tar.gz。   - tar.bz2。 - tar。
	ArchiveFormat *string `json:"archive_format,omitempty"`
}

func (o DownloadArchiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadArchiveRequest struct{}"
	}

	return strings.Join([]string{"DownloadArchiveRequest", string(data)}, " ")
}
