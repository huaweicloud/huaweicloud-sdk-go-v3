package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectProtectedRefsUserGroupsRequest Request Object
type ListProjectProtectedRefsUserGroupsRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[[查询项目列表](https://support.huaweicloud.com/eu/api-projectman/ListProjectsV4.html)](tag:hws_eu)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询关键字，可模糊匹配用户名称、用户昵称、租户名称。
	Search *string `json:"search,omitempty"`
}

func (o ListProjectProtectedRefsUserGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectProtectedRefsUserGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectProtectedRefsUserGroupsRequest", string(data)}, " ")
}
