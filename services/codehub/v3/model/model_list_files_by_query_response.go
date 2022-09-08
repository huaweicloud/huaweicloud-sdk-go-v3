package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFilesByQueryResponse struct {
	Error *Error `json:"error,omitempty"`

	// 差异列表
	Result *[]FileContentInfo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFilesByQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFilesByQueryResponse struct{}"
	}

	return strings.Join([]string{"ListFilesByQueryResponse", string(data)}, " ")
}
