package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteTestCaseResponse struct {
	// 接口调用失败错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 接口调用失败错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorDetail    *ErrorDetailInfo `json:"error_detail,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchDeleteTestCaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCaseResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCaseResponse", string(data)}, " ")
}
