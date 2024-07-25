package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataobjectRelationsResponse Response Object
type CreateDataobjectRelationsResponse struct {

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

	Data *BatchOperateDataobjectResult `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataobjectRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataobjectRelationsResponse struct{}"
	}

	return strings.Join([]string{"CreateDataobjectRelationsResponse", string(data)}, " ")
}
