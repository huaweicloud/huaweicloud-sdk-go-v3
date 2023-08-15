package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoStatistics struct {

	// 添加代码行
	AddLines *int32 `json:"add_lines,omitempty"`

	// 分支名
	Branch *string `json:"branch,omitempty"`

	// 提交次数
	CommitCount *int32 `json:"commit_count,omitempty"`

	// 仓库统计创建的时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 删除代码行
	DeleteLines *int32 `json:"delete_lines,omitempty"`

	// 仓库统计记录id
	Id *int32 `json:"id,omitempty"`

	// 仓库id
	ProjectId *int32 `json:"project_id,omitempty"`

	// 仓库统计更新的时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o RepoStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoStatistics struct{}"
	}

	return strings.Join([]string{"RepoStatistics", string(data)}, " ")
}
