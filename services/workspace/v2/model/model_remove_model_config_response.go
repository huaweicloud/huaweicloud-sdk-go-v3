package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveModelConfigResponse Response Object
type RemoveModelConfigResponse struct {

	// 删除数量。
	DeletedCount *int32 `json:"deleted_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails  *[]ModelConfigFailedItem `json:"failed_details,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o RemoveModelConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveModelConfigResponse struct{}"
	}

	return strings.Join([]string{"RemoveModelConfigResponse", string(data)}, " ")
}
