package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteProvidersResponse Response Object
type BatchDeleteProvidersResponse struct {

	// 成功解绑数量。
	DeletedCount *int32 `json:"deleted_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails  *[]BatchOperationFailedItem `json:"failed_details,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o BatchDeleteProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProvidersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteProvidersResponse", string(data)}, " ")
}
