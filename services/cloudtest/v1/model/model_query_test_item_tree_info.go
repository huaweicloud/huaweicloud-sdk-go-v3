package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTestItemTreeInfo struct {

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 阶段
	Stage *string `json:"stage,omitempty"`

	// 活动
	Activity *string `json:"activity,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 任务uri
	TaskUri *string `json:"task_uri,omitempty"`

	// 用例服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 是否包含用例数
	ContainTotal *bool `json:"contain_total,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 排序类型
	SortType *string `json:"sort_type,omitempty"`

	// 页码
	PageNumber *int32 `json:"page_number,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o QueryTestItemTreeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTestItemTreeInfo struct{}"
	}

	return strings.Join([]string{"QueryTestItemTreeInfo", string(data)}, " ")
}
