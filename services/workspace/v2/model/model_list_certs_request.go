package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertsRequest Request Object
type ListCertsRequest struct {

	// 每页显示的数量。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListCertsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertsRequest struct{}"
	}

	return strings.Join([]string{"ListCertsRequest", string(data)}, " ")
}
