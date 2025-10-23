package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryResourcePermissionsRequest Request Object
type ListRepositoryResourcePermissionsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 权限资源名称 **约束限制：** 不涉及。 **取值范围：** - repository， 仓库 - code，代码 - member，成员 - branch，分支 - tag，Tag - mr，MR - label, 合并请求标签
	ResourceName string `json:"resource_name"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRepositoryResourcePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryResourcePermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryResourcePermissionsRequest", string(data)}, " ")
}
