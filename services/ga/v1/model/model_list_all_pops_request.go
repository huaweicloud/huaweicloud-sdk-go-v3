package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllPopsRequest Request Object
type ListAllPopsRequest struct {

	// 分页查询每页的资源个数。如果不设置，则默认为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询的起始的资源ID，表示上一页最后一条查询资源记录的ID。不指定时表示查询第一页。 必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`
}

func (o ListAllPopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllPopsRequest struct{}"
	}

	return strings.Join([]string{"ListAllPopsRequest", string(data)}, " ")
}
