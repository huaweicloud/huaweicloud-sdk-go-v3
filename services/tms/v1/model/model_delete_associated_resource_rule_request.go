package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssociatedResourceRuleRequest Request Object
type DeleteAssociatedResourceRuleRequest struct {

	// 规则的配置名称。
	SettingName string `json:"setting_name"`

	// 要关闭的规则所在的region集合。
	RegionId []string `json:"region_id"`
}

func (o DeleteAssociatedResourceRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssociatedResourceRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssociatedResourceRuleRequest", string(data)}, " ")
}
