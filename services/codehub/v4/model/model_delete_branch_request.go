package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBranchRequest Request Object
type DeleteBranchRequest struct {

	// **参数解释：** 分支名称。 **取值范围：** 字符串长度不少于1，不超过2000。
	BranchName string `json:"branch_name"`

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`
}

func (o DeleteBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBranchRequest struct{}"
	}

	return strings.Join([]string{"DeleteBranchRequest", string(data)}, " ")
}
