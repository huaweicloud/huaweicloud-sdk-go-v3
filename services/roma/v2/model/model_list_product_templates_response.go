package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProductTemplatesResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 本次返回数量

	Size *int32 `json:"size,omitempty"`
	// 产品模板信息

	Items          *[]ProductTemplate `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListProductTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListProductTemplatesResponse", string(data)}, " ")
}
