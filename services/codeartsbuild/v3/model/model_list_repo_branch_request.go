package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepoBranchRequest Request Object
type ListRepoBranchRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 代码仓库名，不支持中文
	RepositoryName string `json:"repository_name"`
}

func (o ListRepoBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoBranchRequest struct{}"
	}

	return strings.Join([]string{"ListRepoBranchRequest", string(data)}, " ")
}
