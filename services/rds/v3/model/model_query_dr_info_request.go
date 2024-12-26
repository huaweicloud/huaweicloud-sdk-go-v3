package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDrInfoRequest struct {

	// 容灾关系id
	Id *string `json:"id,omitempty"`

	// 容灾搭建状态
	Status *string `json:"status,omitempty"`

	// 灾备实例id
	MasterInstanceId *string `json:"master_instance_id,omitempty"`

	// 主实例region
	MasterRegion *string `json:"master_region,omitempty"`

	// 灾备实例id
	SlaveInstanceId *string `json:"slave_instance_id,omitempty"`

	// 灾备实例region
	SlaveRegion *string `json:"slave_region,omitempty"`

	// 创建起始时间
	CreateAtStart *int64 `json:"create_at_start,omitempty"`

	// 创建结束时间
	CreateAtEnd *int64 `json:"create_at_end,omitempty"`

	// 排序方式。 DESC，降序。 ASC，升序。 默认降序。
	Order *string `json:"order,omitempty"`

	// 排序字段。 status 容灾搭建状态。 time 容灾搭建时间。 master_region 主实例region。 slave_region 灾备实例region 默认容灾搭建时间。
	SortField *string `json:"sort_field,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o QueryDrInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDrInfoRequest struct{}"
	}

	return strings.Join([]string{"QueryDrInfoRequest", string(data)}, " ")
}
