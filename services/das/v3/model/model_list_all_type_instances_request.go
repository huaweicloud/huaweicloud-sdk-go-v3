package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTypeInstancesRequest Request Object
type ListAllTypeInstancesRequest struct {

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`

	// 数据库来源类型
	NetworkType *string `json:"network_type,omitempty"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 实例ID
	Id *string `json:"id,omitempty"`
}

func (o ListAllTypeInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTypeInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAllTypeInstancesRequest", string(data)}, " ")
}
