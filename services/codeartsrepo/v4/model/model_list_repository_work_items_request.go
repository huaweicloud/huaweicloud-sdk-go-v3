package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryWorkItemsRequest Request Object
type ListRepositoryWorkItemsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 是否为IPD类型项目的工作项。 **取值范围：** true/false
	IsIpd bool `json:"is_ipd"`

	// **参数解释：** 工作项标题搜索关键字。 **取值范围：** 字符串长度不少于1，不超过200。
	Subject *string `json:"subject,omitempty"`

	// **参数解释：** 合并请求的目标分支，如果要获取合并请求可关联的工作项列表，则需要传递该字段。  **约束限制：** 不支持以 - . refs/heads/ refs/remotes/ 开头，不支持空格 [ \\ < ~ ^ : ? * ! ( ) ' \" | 等特殊字符，不支持以. / .lock结尾。  **取值范围：** 字符串长度不少于1，不超过200。 **默认取值：** 不涉及。
	TargetBranch *string `json:"target_branch,omitempty"`
}

func (o ListRepositoryWorkItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryWorkItemsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryWorkItemsRequest", string(data)}, " ")
}
