package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTestReportCustomDetailResponse Response Object
type SaveTestReportCustomDetailResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueListTestReportCustomDetailVo `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId *string `json:"request_id,omitempty"`

	// 本次请求的受理的服务地址
	ServerAddress  *string `json:"server_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveTestReportCustomDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTestReportCustomDetailResponse struct{}"
	}

	return strings.Join([]string{"SaveTestReportCustomDetailResponse", string(data)}, " ")
}
