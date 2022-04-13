package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFunctionAsyncInvocationsResponse struct {
	// 异步调用记录列表。

	Invocations    *[]ListFunctionAsyncInvocationsResult `json:"invocations,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListFunctionAsyncInvocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvocationsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvocationsResponse", string(data)}, " ")
}
