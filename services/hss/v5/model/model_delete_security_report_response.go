package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityReportResponse Response Object
type DeleteSecurityReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityReportResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityReportResponse", string(data)}, " ")
}
