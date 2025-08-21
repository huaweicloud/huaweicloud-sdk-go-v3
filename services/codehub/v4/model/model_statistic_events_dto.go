package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticEventsDto struct {

	// **参数解释：** 统计ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户ID。
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释：** 仓库ID。
	ProjectId *int32 `json:"project_id,omitempty"`

	// **参数解释：** 分支名。 **取值范围：** 最小1个字节，最大200字节
	Branch *string `json:"branch,omitempty"`

	// **参数解释：** 统计状态。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 统计时间。
	StatDate *string `json:"stat_date,omitempty"`

	// **参数解释：** 统计创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 统计更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o StatisticEventsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticEventsDto struct{}"
	}

	return strings.Join([]string{"StatisticEventsDto", string(data)}, " ")
}
