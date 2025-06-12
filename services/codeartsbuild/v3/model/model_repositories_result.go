package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RepositoriesResult 结果
type RepositoriesResult struct {

	// 最后一次仓库名称
	Latest *string `json:"latest,omitempty"`

	// 本次任务的构建步骤详情，返回的步骤为页面可见步骤
	Repositories *[]string `json:"repositories,omitempty"`
}

func (o RepositoriesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoriesResult struct{}"
	}

	return strings.Join([]string{"RepositoriesResult", string(data)}, " ")
}
