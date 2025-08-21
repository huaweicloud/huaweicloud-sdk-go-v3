package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMergeableStateOuterRequest Request Object
type ShowMergeableStateOuterRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  合并请求 iid。
	MergeRequestIid int32 `json:"merge_request_iid"`
}

func (o ShowMergeableStateOuterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeableStateOuterRequest struct{}"
	}

	return strings.Join([]string{"ShowMergeableStateOuterRequest", string(data)}, " ")
}
