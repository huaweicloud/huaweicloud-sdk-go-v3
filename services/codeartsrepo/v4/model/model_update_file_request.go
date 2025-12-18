package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFileRequest Request Object
type UpdateFileRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过10000。
	FilePath string `json:"file_path"`

	Body *UpdateFileBodyDto `json:"body,omitempty"`
}

func (o UpdateFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFileRequest struct{}"
	}

	return strings.Join([]string{"UpdateFileRequest", string(data)}, " ")
}
