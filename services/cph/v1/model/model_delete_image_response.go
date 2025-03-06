package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageResponse Response Object
type DeleteImageResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 任务错误码说明。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务错误码。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageResponse", string(data)}, " ")
}
