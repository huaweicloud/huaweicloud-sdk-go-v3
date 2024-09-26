package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResponseBody 列表响应体基类。
type ListResponseBody struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`
}

func (o ListResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResponseBody struct{}"
	}

	return strings.Join([]string{"ListResponseBody", string(data)}, " ")
}
