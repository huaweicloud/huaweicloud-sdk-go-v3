package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSecurityReportResponse Response Object
type AddSecurityReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddSecurityReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSecurityReportResponse struct{}"
	}

	return strings.Join([]string{"AddSecurityReportResponse", string(data)}, " ")
}
