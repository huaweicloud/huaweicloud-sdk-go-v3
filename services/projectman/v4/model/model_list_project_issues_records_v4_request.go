package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectIssuesRecordsV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 偏移量 从0开始,offset是limit的整数倍，limit=10,offset=0,10,20...
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页数量 最小1,最大100
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 变更工作项的时间(查询的起始时间,查询的结束时间)
	OperatedTimeInterval *string `json:"operated_time_interval,omitempty" xml:"operated_time_interval"`
}

func (o ListProjectIssuesRecordsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectIssuesRecordsV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectIssuesRecordsV4Request", string(data)}, " ")
}
