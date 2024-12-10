package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainResourceReportsResponse Response Object
type ShowDomainResourceReportsResponse struct {

	// 统计数据
	Body           *[]ReportData `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowDomainResourceReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainResourceReportsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainResourceReportsResponse", string(data)}, " ")
}
