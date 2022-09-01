package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 代码提交记录 / 分支创建记录 详细信息
type CommitRecordDetail struct {

	// 仓库ID
	RepositoryId *string `json:"repository_id,omitempty" xml:"repository_id"`

	// 分支名称
	BranchName *string `json:"branch_name,omitempty" xml:"branch_name"`

	// commit id
	CommitId *string `json:"commit_id,omitempty" xml:"commit_id"`

	// commit short id
	CommitShortId *string `json:"commit_short_id,omitempty" xml:"commit_short_id"`

	// commit 信息
	CommitMsg *string `json:"commit_msg,omitempty" xml:"commit_msg"`

	// commit URL
	CommitUrl *string `json:"commit_url,omitempty" xml:"commit_url"`

	User *SimpleUser `json:"user,omitempty" xml:"user"`

	// 查询的类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 创建时间
	CreateDate *string `json:"create_date,omitempty" xml:"create_date"`

	// 更新时间
	UpdateDate *string `json:"update_date,omitempty" xml:"update_date"`
}

func (o CommitRecordDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitRecordDetail struct{}"
	}

	return strings.Join([]string{"CommitRecordDetail", string(data)}, " ")
}
