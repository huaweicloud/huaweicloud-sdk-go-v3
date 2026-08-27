package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindProvidersResponse Response Object
type BatchBindProvidersResponse struct {

	// 请求总数。
	Total *int32 `json:"total,omitempty"`

	// 成功绑定数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails  *[]BatchOperationFailedItem `json:"failed_details,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o BatchBindProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindProvidersResponse struct{}"
	}

	return strings.Join([]string{"BatchBindProvidersResponse", string(data)}, " ")
}
