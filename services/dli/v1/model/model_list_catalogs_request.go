package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogsRequest Request Object
type ListCatalogsRequest struct {

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页的单页数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCatalogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogsRequest struct{}"
	}

	return strings.Join([]string{"ListCatalogsRequest", string(data)}, " ")
}
