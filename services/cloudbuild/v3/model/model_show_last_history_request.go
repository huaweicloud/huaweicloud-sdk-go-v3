package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowLastHistoryRequest struct {

	// DevCloud项目ID，32位数字、小写字母组合。[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)
	ProjectId string `json:"project_id"`

	// 代码仓库名，不支持中文
	RepositoryName string `json:"repository_name"`
}

func (o ShowLastHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLastHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowLastHistoryRequest", string(data)}, " ")
}
