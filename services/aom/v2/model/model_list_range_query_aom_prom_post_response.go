package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRangeQueryAomPromPostResponse Response Object
type ListRangeQueryAomPromPostResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 响应数据。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRangeQueryAomPromPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRangeQueryAomPromPostResponse struct{}"
	}

	return strings.Join([]string{"ListRangeQueryAomPromPostResponse", string(data)}, " ")
}
