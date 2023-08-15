package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkitemsRequest Request Object
type ListWorkitemsRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 偏移量 从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量 最小1,最大100
	Limit *int32 `json:"limit,omitempty"`

	// 创建工作项的时间(查询的起始时间,查询的结束时间)
	CreatedTimeInterval *string `json:"created_time_interval,omitempty"`
}

func (o ListWorkitemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkitemsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkitemsRequest", string(data)}, " ")
}
