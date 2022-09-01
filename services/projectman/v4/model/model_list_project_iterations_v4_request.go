package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectIterationsV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	// 更新迭代的时间(查询的起始时间,查询的结束时间)
	UpdatedTimeInterval *string `json:"updated_time_interval,omitempty" xml:"updated_time_interval"`

	// 是否包含被删除的迭代,默认false不包含被删除的迭代
	IncludeDeleted *bool `json:"include_deleted,omitempty" xml:"include_deleted"`
}

func (o ListProjectIterationsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectIterationsV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectIterationsV4Request", string(data)}, " ")
}
