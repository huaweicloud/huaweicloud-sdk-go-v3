package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 生成测试套的请求信息
type CreateTestSuitByRepoFileInfo struct {

	// 要生成的测试套名称
	TestsuiteName string `json:"testsuite_name" xml:"testsuite_name"`

	// 仓库id
	RepositoryId string `json:"repository_id" xml:"repository_id"`

	// 仓库分支
	RepositoryBranch string `json:"repository_branch" xml:"repository_branch"`

	// 仓库中yaml或json文件的相对路径，仅支持swagger 2.0版本的yaml和json文件
	FilePath string `json:"file_path" xml:"file_path"`
}

func (o CreateTestSuitByRepoFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestSuitByRepoFileInfo struct{}"
	}

	return strings.Join([]string{"CreateTestSuitByRepoFileInfo", string(data)}, " ")
}
