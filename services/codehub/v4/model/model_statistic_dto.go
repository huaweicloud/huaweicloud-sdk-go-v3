package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticDto struct {

	// **参数解释：** 统计ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库ID。
	ProjectId *int32 `json:"project_id,omitempty"`

	// **参数解释：** 分支名。 **取值范围：** 最小1个字节，最大200字节
	Branch *string `json:"branch,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 增加行数。
	AddLines *int32 `json:"add_lines,omitempty"`

	// **参数解释：** 删除行数。
	DeleteLines *int32 `json:"delete_lines,omitempty"`

	// **参数解释：** 总的提交数量。
	CommitCount *int32 `json:"commit_count,omitempty"`

	// **参数解释：** 最早提交时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 最新更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o StatisticDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticDto struct{}"
	}

	return strings.Join([]string{"StatisticDto", string(data)}, " ")
}
