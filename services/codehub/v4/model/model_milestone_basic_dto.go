package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MilestoneBasicDto struct {

	// 里程碑id
	Id *int32 `json:"id,omitempty"`

	// 里程碑iid
	Iid *int32 `json:"iid,omitempty"`

	// 仓库id
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 里程碑标题
	Title *string `json:"title,omitempty"`

	// 里程碑描述
	Description *string `json:"description,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 到期时间
	DueDate *string `json:"due_date,omitempty"`

	// 开始时间
	StartDate *string `json:"start_date,omitempty"`

	// 仓库路径
	RepositoryPath *string `json:"repository_path,omitempty"`

	// 主页url
	WebUrl *string `json:"web_url,omitempty"`
}

func (o MilestoneBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MilestoneBasicDto struct{}"
	}

	return strings.Join([]string{"MilestoneBasicDto", string(data)}, " ")
}
