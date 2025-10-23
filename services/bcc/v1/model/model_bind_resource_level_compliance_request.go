package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindResourceLevelComplianceRequest Request Object
type BindResourceLevelComplianceRequest struct {

	// 账户ID
	DomainId string `json:"domain_id"`

	// 资源等级ID
	LevelId string `json:"level_id"`

	Body *BindComplianceBody `json:"body,omitempty"`
}

func (o BindResourceLevelComplianceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindResourceLevelComplianceRequest struct{}"
	}

	return strings.Join([]string{"BindResourceLevelComplianceRequest", string(data)}, " ")
}
