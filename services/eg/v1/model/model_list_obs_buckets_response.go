package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketsResponse Response Object
type ListObsBucketsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]ObsBucketInfo `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListObsBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListObsBucketsResponse", string(data)}, " ")
}
