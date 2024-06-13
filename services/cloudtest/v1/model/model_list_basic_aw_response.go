package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasicAwResponse Response Object
type ListBasicAwResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorOfApiTest `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	Result *BasicAwRes `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBasicAwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasicAwResponse struct{}"
	}

	return strings.Join([]string{"ListBasicAwResponse", string(data)}, " ")
}
