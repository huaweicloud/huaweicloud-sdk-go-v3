package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAsyncInvocationsResponse Response Object
type ListAsyncInvocationsResponse struct {

	// 异步调用记录列表。
	Invocations *[]ListFunctionAsyncInvocationsResult `json:"invocations,omitempty"`

	// 查询数据总条数
	Count *int32 `json:"count,omitempty"`

	// 查询下一页的起始位置
	NextMarker     *int32 `json:"next_marker,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAsyncInvocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncInvocationsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncInvocationsResponse", string(data)}, " ")
}
