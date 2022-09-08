package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMonthUsagesResponse struct {

	// 月用量列表
	MonthUsages    *[]MonthUsageVo `json:"month_usages,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowMonthUsagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthUsagesResponse struct{}"
	}

	return strings.Join([]string{"ShowMonthUsagesResponse", string(data)}, " ")
}
