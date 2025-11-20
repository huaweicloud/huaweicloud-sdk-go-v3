package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateForwardRuleRequestBody struct {

	// 源站IP
	SourceIp string `json:"source_ip"`
}

func (o UpdateForwardRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateForwardRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateForwardRuleRequestBody", string(data)}, " ")
}
