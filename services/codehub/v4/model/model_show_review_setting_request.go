package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviewSettingRequest Request Object
type ShowReviewSettingRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 额外返回可勾选检视意见分类和系统预置检视意见分类。 **取值范围：** - true, 返回可勾选检视意见分类和系统预置检视意见分类。 - false, 不返回可勾选检视意见分类和系统预置检视意见分类。
	WithDefaultReviewCategories *bool `json:"with_default_review_categories,omitempty"`
}

func (o ShowReviewSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviewSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowReviewSettingRequest", string(data)}, " ")
}
