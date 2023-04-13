package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowReportResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueCustomReportListVo `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportResponse struct{}"
	}

	return strings.Join([]string{"ShowReportResponse", string(data)}, " ")
}
