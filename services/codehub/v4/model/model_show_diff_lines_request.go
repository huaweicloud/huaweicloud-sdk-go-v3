package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiffLinesRequest Request Object
type ShowDiffLinesRequest struct {

	// **参数解释：** 仓库的ID，通过[[查询用户所有仓库](https://support.huaweicloud.com/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws)[[查询用户所有仓库](https://support.huaweicloud.com/intl/en-us/api-codeartsrepo/ListUserAllRepositories.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **约束限制：** 不涉及。
	RepositoryId int32 `json:"repository_id"`

	// **参数解释：** 文件路径。 **取值范围：** 字符串长度不少于1，不超过10000。
	FilePath string `json:"file_path"`

	// **参数解释：** commit ID。 **取值范围：** 字符串长度不少于1，不超过40。
	CommitId string `json:"commit_id"`

	// **参数解释：** 文件开始行数。
	Start int32 `json:"start"`

	// **参数解释：** 文件结束行数。
	End int32 `json:"end"`
}

func (o ShowDiffLinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffLinesRequest struct{}"
	}

	return strings.Join([]string{"ShowDiffLinesRequest", string(data)}, " ")
}
