package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListValueListResponse struct {

	// 引用表条数
	Total *int32 `json:"total,omitempty"`

	// 引用表列表
	Items          *[]ValueListResponseBody `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListValueListResponse struct{}"
	}

	return strings.Join([]string{"ListValueListResponse", string(data)}, " ")
}
