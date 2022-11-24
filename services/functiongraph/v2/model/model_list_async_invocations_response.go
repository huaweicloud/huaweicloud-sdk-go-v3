package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAsyncInvocationsResponse struct {

	// 异步调用记录列表。
	Invocations    *[]ListFunctionAsyncInvocationsResult `json:"invocations,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListAsyncInvocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncInvocationsResponse struct{}"
	}

	return strings.Join([]string{"ListAsyncInvocationsResponse", string(data)}, " ")
}
