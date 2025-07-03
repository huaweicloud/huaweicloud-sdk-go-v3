package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTicketParamsV2CountFilters struct {

	// 过滤名称
	Name *string `json:"name,omitempty"`

	// 批量计数
	Filters *[]ObjectFilterV2 `json:"filters,omitempty"`
}

func (o ListTicketParamsV2CountFilters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTicketParamsV2CountFilters struct{}"
	}

	return strings.Join([]string{"ListTicketParamsV2CountFilters", string(data)}, " ")
}
