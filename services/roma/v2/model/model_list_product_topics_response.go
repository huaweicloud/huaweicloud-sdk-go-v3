package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductTopicsResponse Response Object
type ListProductTopicsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 产品主题列表
	Items          *[]ProductTopic `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListProductTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListProductTopicsResponse", string(data)}, " ")
}
