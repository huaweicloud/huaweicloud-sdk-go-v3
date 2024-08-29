package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobOperationLogResponse Response Object
type ListJobOperationLogResponse struct {

	// 满足查询要求的操作日志总数
	Count *int32 `json:"count,omitempty"`

	// 操作
	Operations     *[]interface{} `json:"operations,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListJobOperationLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobOperationLogResponse struct{}"
	}

	return strings.Join([]string{"ListJobOperationLogResponse", string(data)}, " ")
}
