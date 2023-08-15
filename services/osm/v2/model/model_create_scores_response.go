package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScoresResponse Response Object
type CreateScoresResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScoresResponse struct{}"
	}

	return strings.Join([]string{"CreateScoresResponse", string(data)}, " ")
}
