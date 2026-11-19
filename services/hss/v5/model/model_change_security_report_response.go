package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityReportResponse Response Object
type ChangeSecurityReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeSecurityReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityReportResponse struct{}"
	}

	return strings.Join([]string{"ChangeSecurityReportResponse", string(data)}, " ")
}
