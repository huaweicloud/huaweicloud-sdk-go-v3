package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListByoipPoolsRequest Request Object
type ListByoipPoolsRequest struct {

	// 分页查询每页的资源个数。如果不设置，则默认为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源ID，表示上一页最后一条查询资源记录的ID。不指定时表示查询第一页。 必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`
}

func (o ListByoipPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListByoipPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListByoipPoolsRequest", string(data)}, " ")
}
