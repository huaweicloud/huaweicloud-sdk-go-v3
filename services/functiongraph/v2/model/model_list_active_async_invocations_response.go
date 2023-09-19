package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActiveAsyncInvocationsResponse Response Object
type ListActiveAsyncInvocationsResponse struct {

	// 异步调用记录列表。
	Invocations    *[]ListFunctionAsyncInvocationsResult `json:"invocations,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListActiveAsyncInvocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveAsyncInvocationsResponse struct{}"
	}

	return strings.Join([]string{"ListActiveAsyncInvocationsResponse", string(data)}, " ")
}
