package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTestReportResponse Response Object
type BatchDeleteTestReportResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchDeleteTestReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTestReportResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestReportResponse", string(data)}, " ")
}
