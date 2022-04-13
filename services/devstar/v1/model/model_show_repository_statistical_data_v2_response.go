package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepositoryStatisticalDataV2Response struct {
	// 代码仓的名称

	Name *string `json:"name,omitempty"`
	// 提交数量

	CommitNumber *int32 `json:"commit_number,omitempty"`
	// Git库容量

	GitRepoCap *string `json:"git_repo_cap,omitempty"`
	// 近一次提交时间

	LastCommitTime *string `json:"last_commit_time,omitempty"`
	// 代码行数

	CodeLines *int32 `json:"code_lines,omitempty"`
	// 分支数量

	BranchNumber *int32 `json:"branch_number,omitempty"`
	// 代码仓路径url

	DetailUrl *string `json:"detail_url,omitempty"`
	// 代码仓下载url

	DownloadUrl    *string `json:"download_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryStatisticalDataV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticalDataV2Response struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticalDataV2Response", string(data)}, " ")
}
