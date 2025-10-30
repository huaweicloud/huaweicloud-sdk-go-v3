package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendSecurityReportResponse Response Object
type SendSecurityReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendSecurityReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendSecurityReportResponse struct{}"
	}

	return strings.Join([]string{"SendSecurityReportResponse", string(data)}, " ")
}
