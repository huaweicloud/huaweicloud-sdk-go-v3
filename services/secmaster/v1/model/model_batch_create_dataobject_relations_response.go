package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDataobjectRelationsResponse Response Object
type BatchCreateDataobjectRelationsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	Data           *BatchOperateDataobjectResult `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o BatchCreateDataobjectRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDataobjectRelationsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDataobjectRelationsResponse", string(data)}, " ")
}
