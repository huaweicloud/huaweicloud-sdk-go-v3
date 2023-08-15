package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllStandardsRequest Request Object
type ListAllStandardsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 目录ID。获取该目录下的数据，如果有子目录，获取所有子目录的数据
	DirectoryId *string `json:"directory_id,omitempty"`

	// 时间过滤左边界,与end_time一起使用,只支持时间范围过滤,单边过滤无效
	BeginTime *string `json:"begin_time,omitempty"`

	// 时间过滤右边界,与begin_time一起使用只支持时间范围过滤,单边过滤无效
	EndTime *string `json:"end_time,omitempty"`

	// 查询条数，即查询Y条数据。默认值50，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整。默认值0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAllStandardsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllStandardsRequest struct{}"
	}

	return strings.Join([]string{"ListAllStandardsRequest", string(data)}, " ")
}
