package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageMemberResponse Response Object
type DeleteImageMemberResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 任务错误码说明。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务错误码。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageMemberResponse", string(data)}, " ")
}
