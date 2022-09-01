package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoStatisticsEvent struct {

	// 分支名
	Branch *string `json:"branch,omitempty" xml:"branch"`

	// 仓库统计创建的时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty" xml:"created_at"`

	// 仓库统计的日期
	Date *string `json:"date,omitempty" xml:"date"`

	// 仓库统计事件的id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 仓库id
	ProjectId *int32 `json:"project_id,omitempty" xml:"project_id"`

	// 仓库统计的状态: 等待统计waiting  正在统计active  完成统计finish
	Status *string `json:"status,omitempty" xml:"status"`

	// 仓库统计更新的时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty" xml:"updated_at"`

	// 用户id
	UserId *int32 `json:"user_id,omitempty" xml:"user_id"`
}

func (o RepoStatisticsEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoStatisticsEvent struct{}"
	}

	return strings.Join([]string{"RepoStatisticsEvent", string(data)}, " ")
}
