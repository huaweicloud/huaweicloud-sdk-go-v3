package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeInstanceReportNewResponse Response Object
type UnsubscribeInstanceReportNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnsubscribeInstanceReportNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeInstanceReportNewResponse struct{}"
	}

	return strings.Join([]string{"UnsubscribeInstanceReportNewResponse", string(data)}, " ")
}
