package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDailyReportRequest struct {
	// 用户EIP对应的ID

	FloatingIpId string `json:"floating_ip_id"`
	// 用户EIP

	Ip *string `json:"ip,omitempty"`
}

func (o ListDailyReportRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDailyReportRequest struct{}"
	}

	return strings.Join([]string{"ListDailyReportRequest", string(data)}, " ")
}
