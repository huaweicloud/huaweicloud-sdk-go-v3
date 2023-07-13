package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitRecordDetail 代码提交记录 / 分支创建记录 详细信息
type CommitRecordDetail struct {

	// 仓库ID
	RepositoryId *string `json:"repository_id,omitempty"`

	// 分支名称
	BranchName *string `json:"branch_name,omitempty"`

	// commit id
	CommitId *string `json:"commit_id,omitempty"`

	// commit short id
	CommitShortId *string `json:"commit_short_id,omitempty"`

	// commit 信息
	CommitMsg *string `json:"commit_msg,omitempty"`

	// commit URL
	CommitUrl *string `json:"commit_url,omitempty"`

	User *SimpleUser `json:"user,omitempty"`

	// 查询的类型
	Type *string `json:"type,omitempty"`

	// 创建时间
	CreateDate *string `json:"create_date,omitempty"`

	// 更新时间
	UpdateDate *string `json:"update_date,omitempty"`
}

func (o CommitRecordDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitRecordDetail struct{}"
	}

	return strings.Join([]string{"CommitRecordDetail", string(data)}, " ")
}
