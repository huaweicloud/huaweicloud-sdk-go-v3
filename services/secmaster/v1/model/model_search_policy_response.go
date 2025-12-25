package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPolicyResponse Response Object
type SearchPolicyResponse struct {

	// 返回编码
	Code *string `json:"code,omitempty"`

	// 返回数据
	Data *[]interface{} `json:"data,omitempty"`

	// 返回消息
	Message *string `json:"message,omitempty"`

	// 页码
	Page *int32 `json:"page,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 单页数量
	Size *int32 `json:"size,omitempty"`

	// 成功状态
	Success *bool `json:"success,omitempty"`

	// 总量
	Total *int32 `json:"total,omitempty"`

	ContentType    *string `json:"content-type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyResponse struct{}"
	}

	return strings.Join([]string{"SearchPolicyResponse", string(data)}, " ")
}
