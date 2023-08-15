package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepositoryResponse Response Object
type DeleteRepositoryResponse struct {
	Error *Error `json:"error,omitempty"`

	// 响应结果
	Result *bool `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepositoryResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepositoryResponse", string(data)}, " ")
}
