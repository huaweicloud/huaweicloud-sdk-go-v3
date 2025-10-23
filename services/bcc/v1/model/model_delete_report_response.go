package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReportResponse Response Object
type DeleteReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReportResponse struct{}"
	}

	return strings.Join([]string{"DeleteReportResponse", string(data)}, " ")
}
