package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindResourceLevelComplianceRequest Request Object
type UnbindResourceLevelComplianceRequest struct {

	// 账户ID
	DomainId string `json:"domain_id"`

	// 资源等级ID
	LevelId string `json:"level_id"`
}

func (o UnbindResourceLevelComplianceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindResourceLevelComplianceRequest struct{}"
	}

	return strings.Join([]string{"UnbindResourceLevelComplianceRequest", string(data)}, " ")
}
