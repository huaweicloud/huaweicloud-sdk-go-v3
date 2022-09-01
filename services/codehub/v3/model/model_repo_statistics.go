package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoStatistics struct {

	// 添加代码行
	AddLines *int32 `json:"add_lines,omitempty" xml:"add_lines"`

	// 分支名
	Branch *string `json:"branch,omitempty" xml:"branch"`

	// 提交次数
	CommitCount *int32 `json:"commit_count,omitempty" xml:"commit_count"`

	// 仓库统计创建的时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty" xml:"created_at"`

	// 删除代码行
	DeleteLines *int32 `json:"delete_lines,omitempty" xml:"delete_lines"`

	// 仓库id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 仓库id
	ProjectId *int32 `json:"project_id,omitempty" xml:"project_id"`

	// 仓库统计更新的时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty" xml:"updated_at"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`
}

func (o RepoStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoStatistics struct{}"
	}

	return strings.Join([]string{"RepoStatistics", string(data)}, " ")
}
