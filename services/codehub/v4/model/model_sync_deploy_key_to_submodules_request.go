package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncDeployKeyToSubmodulesRequest Request Object
type SyncDeployKeyToSubmodulesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：**  部署密钥 ID。
	KeyId int32 `json:"key_id"`
}

func (o SyncDeployKeyToSubmodulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncDeployKeyToSubmodulesRequest struct{}"
	}

	return strings.Join([]string{"SyncDeployKeyToSubmodulesRequest", string(data)}, " ")
}
