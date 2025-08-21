package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProtectedTagRequest Request Object
type UpdateProtectedTagRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 保护分Tag或通配符列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TagName string `json:"tag_name"`

	// 保护Tag权限列表
	Body *[]RepositoryProtectedTagActionBodyDto `json:"body,omitempty"`
}

func (o UpdateProtectedTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedTagRequest struct{}"
	}

	return strings.Join([]string{"UpdateProtectedTagRequest", string(data)}, " ")
}
