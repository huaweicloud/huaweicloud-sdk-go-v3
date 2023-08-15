package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFilesByQueryResponse Response Object
type ListFilesByQueryResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *FileContentInfo `json:"result,omitempty"`

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
