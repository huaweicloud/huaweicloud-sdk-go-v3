package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCgwsRequest Request Object
type ListCgwsRequest struct {

	// CustomerGateway ID
	CgwId *string `json:"cgw_id,omitempty"`

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的创建时间，为空时为查询第一页。使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`
}

func (o ListCgwsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCgwsRequest struct{}"
	}

	return strings.Join([]string{"ListCgwsRequest", string(data)}, " ")
}
