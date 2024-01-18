package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppConnectionRequest Request Object
type ListAppConnectionRequest struct {

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	Body *ListAppConnectionReq `json:"body,omitempty"`
}

func (o ListAppConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppConnectionRequest struct{}"
	}

	return strings.Join([]string{"ListAppConnectionRequest", string(data)}, " ")
}
