package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFirewallReportResponse Response Object
type ShowFirewallReportResponse struct {
	Data           *FirewallReport `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowFirewallReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFirewallReportResponse struct{}"
	}

	return strings.Join([]string{"ShowFirewallReportResponse", string(data)}, " ")
}
