package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueriesResponse Response Object
type ListQueriesResponse struct {

	// 响应码。
	Code *int32 `json:"code,omitempty"`

	// 响应信息。
	Msg *string `json:"msg,omitempty"`

	Data *ListQueriesData `json:"data,omitempty"`

	// 总条数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListQueriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesResponse struct{}"
	}

	return strings.Join([]string{"ListQueriesResponse", string(data)}, " ")
}
