package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoStatisticsEvent struct {
	// 分支名

	Branch *string `json:"branch,omitempty"`
	// 仓库统计创建的时间

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 仓库统计的日期

	Date *string `json:"date,omitempty"`
	// 仓库统计事件的id

	Id *int32 `json:"id,omitempty"`
	// 仓库id

	ProjectId *int32 `json:"project_id,omitempty"`
	// 仓库统计的状态: 等待统计waiting  正在统计active  完成统计finish

	Status *string `json:"status,omitempty"`
	// 仓库统计更新的时间

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
	// 用户id

	UserId *int32 `json:"user_id,omitempty"`
}

func (o RepoStatisticsEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoStatisticsEvent struct{}"
	}

	return strings.Join([]string{"RepoStatisticsEvent", string(data)}, " ")
}
