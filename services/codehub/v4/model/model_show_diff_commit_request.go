package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiffCommitRequest Request Object
type ShowDiffCommitRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 分支名、tag名、提交ID。
	Sha string `json:"sha"`

	// **参数解释：** 是否忽略空白数量更改。 **取值范围：** - true，忽略空白数量的更改。 - false，不会忽略空白数量的更改。
	IgnoreWhitespaceChange *bool `json:"ignore_whitespace_change,omitempty"`

	// **参数解释：** 是否返回统计数量。 **取值范围：** - true，不返回统计数量。 - false，返回统计数量。
	NotStatistic *bool `json:"not_statistic,omitempty"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDiffCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowDiffCommitRequest", string(data)}, " ")
}
