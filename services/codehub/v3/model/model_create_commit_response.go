package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCommitResponse struct {
	Error *Error `json:"error,omitempty"`
	// 差异列表

	Result *[]CreateCommitResponseBody `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitResponse struct{}"
	}

	return strings.Join([]string{"CreateCommitResponse", string(data)}, " ")
}
