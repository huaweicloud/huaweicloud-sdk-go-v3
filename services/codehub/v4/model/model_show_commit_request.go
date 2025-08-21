package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitRequest Request Object
type ShowCommitRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支名、tag名、提交ID。
	Sha string `json:"sha"`

	// **参数解释：** 是否包含状态信息。 **取值范围：** - true，包含。 - false，不包含。
	Stats *bool `json:"stats,omitempty"`

	// **参数解释：** 是否包含代码变化信息。 **取值范围：** - true，包含。 - false，不包含。
	ShowCodeChanges *bool `json:"show_code_changes,omitempty"`
}

func (o ShowCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowCommitRequest", string(data)}, " ")
}
