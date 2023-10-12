package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationConformancePackRequest Request Object
type ShowOrganizationConformancePackRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`
}

func (o ShowOrganizationConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationConformancePackRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationConformancePackRequest", string(data)}, " ")
}
