package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionsResponse Response Object
type ListExecutionsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 返回数据
	Data *[]ListExecutionResponseData `json:"data,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionsResponse struct{}"
	}

	return strings.Join([]string{"ListExecutionsResponse", string(data)}, " ")
}
