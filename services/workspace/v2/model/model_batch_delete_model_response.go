package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteModelResponse Response Object
type BatchDeleteModelResponse struct {

	// 删除数量。
	DeletedCount *int32 `json:"deleted_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails  *[]ModelBatchDeleteRespFailedDetails `json:"failed_details,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o BatchDeleteModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteModelResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteModelResponse", string(data)}, " ")
}
