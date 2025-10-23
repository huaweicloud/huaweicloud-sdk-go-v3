package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryE2eSettingRequest Request Object
type ShowRepositoryE2eSettingRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 是否查询生效的E2E配置。 **取值范围：** true：查询仓库实际生效的E2E配置，比如继承自代码组或者项目的E2E设置，false：查询仓库自身的E2E配置。
	TakeEffect *bool `json:"take_effect,omitempty"`
}

func (o ShowRepositoryE2eSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryE2eSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowRepositoryE2eSettingRequest", string(data)}, " ")
}
