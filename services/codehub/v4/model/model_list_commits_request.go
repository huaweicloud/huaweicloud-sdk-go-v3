package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommitsRequest Request Object
type ListCommitsRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分支/tag名
	RefName *string `json:"ref_name,omitempty"`

	// 起始时间（不包含）
	Since *string `json:"since,omitempty"`

	// 结束时间（不包含）
	Until *string `json:"until,omitempty"`

	// The file path
	Path *string `json:"path,omitempty"`

	// 提交信息或者commit id
	Message *string `json:"message,omitempty"`

	// 代码作者名称
	Author *string `json:"author,omitempty"`

	// 是否按照时间降序排序
	OrderByDate *bool `json:"order_by_date,omitempty"`

	// 为true时可以追踪文件移动、改名前后的提交记录
	Follow *bool `json:"follow,omitempty"`
}

func (o ListCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListCommitsRequest", string(data)}, " ")
}
