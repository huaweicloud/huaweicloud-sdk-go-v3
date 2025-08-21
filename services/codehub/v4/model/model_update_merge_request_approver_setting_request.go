package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestApproverSettingRequest Request Object
type UpdateMergeRequestApproverSettingRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 合并请求审核设置id。
	SettingId int32 `json:"setting_id"`

	Body *MergeRequestApproverSettingResultDto `json:"body,omitempty"`
}

func (o UpdateMergeRequestApproverSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestApproverSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestApproverSettingRequest", string(data)}, " ")
}
