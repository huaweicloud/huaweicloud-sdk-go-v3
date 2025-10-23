package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindComplianceBody struct {

	// 合规规则id
	ComplianceId string `json:"compliance_id"`
}

func (o BindComplianceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindComplianceBody struct{}"
	}

	return strings.Join([]string{"BindComplianceBody", string(data)}, " ")
}
