package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportsResponse Response Object
type ListReportsResponse struct {

	// 对外时：success|error; 对内时：ok|failed
	Status *string `json:"status,omitempty"`

	Result         *ResultValueListCustomReportListVo `json:"result,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportsResponse struct{}"
	}

	return strings.Join([]string{"ListReportsResponse", string(data)}, " ")
}
