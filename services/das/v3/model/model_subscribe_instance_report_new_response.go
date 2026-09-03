package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeInstanceReportNewResponse Response Object
type SubscribeInstanceReportNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SubscribeInstanceReportNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeInstanceReportNewResponse struct{}"
	}

	return strings.Join([]string{"SubscribeInstanceReportNewResponse", string(data)}, " ")
}
