package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupTreeResponse Response Object
type ListGroupTreeResponse struct {

	// 返回结果
	Result *[]JobGroupTreeResponseBody `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGroupTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupTreeResponse struct{}"
	}

	return strings.Join([]string{"ListGroupTreeResponse", string(data)}, " ")
}
