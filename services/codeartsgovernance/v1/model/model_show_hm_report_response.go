package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHmReportResponse Response Object
type ShowHmReportResponse struct {

	// 响应码
	Code *string `json:"code,omitempty"`

	Data           *HmVulnInfoData `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHmReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHmReportResponse struct{}"
	}

	return strings.Join([]string{"ShowHmReportResponse", string(data)}, " ")
}
