package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVgwsRequest Request Object
type ListP2cVgwsRequest struct {

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的id，为空时为查询第一页。使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`
}

func (o ListP2cVgwsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwsRequest struct{}"
	}

	return strings.Join([]string{"ListP2cVgwsRequest", string(data)}, " ")
}
