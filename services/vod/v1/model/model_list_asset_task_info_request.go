package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetTaskInfoRequest Request Object
type ListAssetTaskInfoRequest struct {

	// 任务类型
	Type *string `json:"type,omitempty"`

	// 媒资Id
	AssetId *string `json:"asset_id,omitempty"`

	// 根据任务创建时间匹配该时间之后的，包含该时间点，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z
	CreateTimeAfter *string `json:"create_time_after,omitempty"`

	// 根据任务创建时间匹配该时间之前的，不包含该时间点，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z
	CreateTimeBefore *string `json:"create_time_before,omitempty"`

	// 根据任务结束时间匹配该时间之后的，包含该时间点，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z
	EndTimeAfter *string `json:"end_time_after,omitempty"`

	// 根据任务结束时间匹配该时间之前的，不包含该时间点，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z
	EndTimeBefore *string `json:"end_time_before,omitempty"`

	// 任务状态
	Status *[]string `json:"status,omitempty"`

	// 标志位，不传默认表示从第一条开始
	Marker *string `json:"marker,omitempty"`

	// 返回的每页个数，默认10，最大100，最小1
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAssetTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ListAssetTaskInfoRequest", string(data)}, " ")
}
