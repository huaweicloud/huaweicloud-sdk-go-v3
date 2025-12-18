package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepositoryResourcePermissionsRequest Request Object
type UpdateRepositoryResourcePermissionsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[[查询用户所有仓库](https://support.huaweicloud.com/eu/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 权限资源名称 **约束限制：** 不涉及。 **取值范围：** - repository， 仓库 - code，代码 - member，成员 - branch，分支 - tag，Tag - mr，MR - label, 合并请求标签
	ResourceName string `json:"resource_name"`

	Body *UpdatePermissionBodyDto `json:"body,omitempty"`
}

func (o UpdateRepositoryResourcePermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryResourcePermissionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryResourcePermissionsRequest", string(data)}, " ")
}
