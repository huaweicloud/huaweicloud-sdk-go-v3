package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTestCaseResultResponse struct {

	// 接口调用失败错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 接口调用失败错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	ErrorDetail    *ErrorDetailInfo `json:"error_detail,omitempty" xml:"error_detail"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateTestCaseResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseResultResponse struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseResultResponse", string(data)}, " ")
}
