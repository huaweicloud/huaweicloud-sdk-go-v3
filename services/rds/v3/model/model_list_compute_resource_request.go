package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComputeResourceRequest Request Object
type ListComputeResourceRequest struct {

	// 分页查询，大小，默认为10
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询，偏移量，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 引擎名称： mysql、sqlserver、postgresql
	Engine *string `json:"engine,omitempty"`
}

func (o ListComputeResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputeResourceRequest struct{}"
	}

	return strings.Join([]string{"ListComputeResourceRequest", string(data)}, " ")
}
