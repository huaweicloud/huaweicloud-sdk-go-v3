package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFileRequest Request Object
type DeleteFileRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过10000。
	FilePath string `json:"file_path"`

	// **参数解释：** 作者名称
	AuthorName *string `json:"author_name,omitempty"`

	// **参数解释：** 分支名称。 **取值范围：** 字符串长度不少于1，不超过2000。
	Branch string `json:"branch"`

	// **参数解释：** 删除描述。 **取值范围：** 字符串长度不少于1，不超过2000。
	CommitMessage string `json:"commit_message"`

	// **参数解释：** 作者邮箱。 **取值范围：** 字符串长度不少于1，不超过2000。
	AuthorEmail *string `json:"author_email,omitempty"`
}

func (o DeleteFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteFileRequest", string(data)}, " ")
}
