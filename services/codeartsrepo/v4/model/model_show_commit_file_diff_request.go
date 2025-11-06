package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitFileDiffRequest Request Object
type ShowCommitFileDiffRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支名、tag名、提交ID。
	Sha string `json:"sha"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过100000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 改名之前的文件路径。 **取值范围：** 字符串长度不少于1，不超过100000。
	OldPath *string `json:"old_path,omitempty"`

	// **参数解释：** 是否忽略空白数量更改。 **取值范围：** - true，忽略空白数量的更改。 - false，不会忽略空白数量的更改。
	IgnoreWhitespaceChange *bool `json:"ignore_whitespace_change,omitempty"`
}

func (o ShowCommitFileDiffRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitFileDiffRequest struct{}"
	}

	return strings.Join([]string{"ShowCommitFileDiffRequest", string(data)}, " ")
}
