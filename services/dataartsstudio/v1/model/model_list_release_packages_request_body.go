package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListReleasePackagesRequestBody struct {

	// 包名package_name关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 申请人名称
	ApplyUserName *string `json:"apply_user_name,omitempty"`

	// 发布人名称
	DeployUserName *string `json:"deploy_user_name,omitempty"`

	// 申请开始时间，13位时间戳
	ApplyBeginTime *int64 `json:"apply_begin_time,omitempty"`

	// 申请结束时间，13位时间戳
	ApplyEndTime *int64 `json:"apply_end_time,omitempty"`

	// 发布开始时间，13位时间戳
	DeployBeginTime *int64 `json:"deploy_begin_time,omitempty"`

	// 发布结束时间，13位时间戳
	DeployEndTime *int64 `json:"deploy_end_time,omitempty"`

	// 申请人名称集合，根据该字段筛选，如果选择了apply_user_name，则该名称必须包含在集合内
	ApplyUserNameFilter *[]string `json:"apply_user_name_filter,omitempty"`

	// 发布人名称集合，根据该字段筛选，如果选择了apply_user_name，则该名称必须包含在集合内
	DeployUserNameFilter *[]string `json:"deploy_user_name_filter,omitempty"`

	// 发布状态集合: 1:待审批,2:成功,3:失败, 5:发布中
	DeployStatusFilter *[]int32 `json:"deploy_status_filter,omitempty"`

	// 排序方向，默认是desc
	SortedDirection *string `json:"sorted_direction,omitempty"`

	// 排序字段，默认是apply_timestamp
	OrderColumn *string `json:"order_column,omitempty"`

	// 分页返回结果，默认是10
	Limit *int32 `json:"limit,omitempty"`

	// 分页的起始页，默认值位0，取值范围大于等于0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListReleasePackagesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReleasePackagesRequestBody struct{}"
	}

	return strings.Join([]string{"ListReleasePackagesRequestBody", string(data)}, " ")
}
