package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRangeQueryAomPromGetResponse Response Object
type ListRangeQueryAomPromGetResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 响应数据。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRangeQueryAomPromGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRangeQueryAomPromGetResponse struct{}"
	}

	return strings.Join([]string{"ListRangeQueryAomPromGetResponse", string(data)}, " ")
}
