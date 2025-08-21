package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCommitRevertRequest Request Object
type CreateCommitRevertRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 回退的CommitId **取值范围：** 字符串长度40。
	Sha string `json:"sha"`

	Body *CommitRevetOrCherryPickDto `json:"body,omitempty"`
}

func (o CreateCommitRevertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitRevertRequest struct{}"
	}

	return strings.Join([]string{"CreateCommitRevertRequest", string(data)}, " ")
}
