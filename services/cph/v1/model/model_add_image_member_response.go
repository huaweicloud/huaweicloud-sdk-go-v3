package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageMemberResponse Response Object
type AddImageMemberResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 任务错误码说明。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务错误码。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageMemberResponse struct{}"
	}

	return strings.Join([]string{"AddImageMemberResponse", string(data)}, " ")
}
