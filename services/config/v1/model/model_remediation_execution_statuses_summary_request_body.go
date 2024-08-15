package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationExecutionStatusesSummaryRequestBody 合规规则最新修正记录查询条件。
type RemediationExecutionStatusesSummaryRequestBody struct {

	// 合规规则最新修正记录查询条件列表。
	ResourceKeys []RemediationResourceKey `json:"resource_keys"`
}

func (o RemediationExecutionStatusesSummaryRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationExecutionStatusesSummaryRequestBody struct{}"
	}

	return strings.Join([]string{"RemediationExecutionStatusesSummaryRequestBody", string(data)}, " ")
}
