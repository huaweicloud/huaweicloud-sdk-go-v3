package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyModelConfigResponse Response Object
type ApplyModelConfigResponse struct {

	// 请求总数。
	Total *int32 `json:"total,omitempty"`

	// 成功数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails  *[]ModelConfigFailedItem `json:"failed_details,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ApplyModelConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyModelConfigResponse struct{}"
	}

	return strings.Join([]string{"ApplyModelConfigResponse", string(data)}, " ")
}
