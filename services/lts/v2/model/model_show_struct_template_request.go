package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowStructTemplateRequest struct {
	// 日志组ID

	LogGroupId string `json:"logGroupId"`
	// 日志流ID

	LogStreamId string `json:"logStreamId"`
}

func (o ShowStructTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateRequest", string(data)}, " ")
}
