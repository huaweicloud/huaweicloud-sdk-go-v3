package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionInstancesResponse Response Object
type ListExecutionInstancesResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 批次实例列表
	Data *[]interface{} `json:"data,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListExecutionInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListExecutionInstancesResponse", string(data)}, " ")
}
