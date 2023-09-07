package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllConnectionsRequest Request Object
type ShowAllConnectionsRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于1
	Offset int32 `json:"offset"`

	// 每页显示的条目数量
	Limit int32 `json:"limit"`

	// 模糊查询参数
	Name *string `json:"name,omitempty"`
}

func (o ShowAllConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ShowAllConnectionsRequest", string(data)}, " ")
}
