package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVaultSummaryRequest Request Object
type ShowVaultSummaryRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域regionId
	RegionId *string `json:"region_id,omitempty"`

	// 企业项目Id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 开始时间,格式为“yyyy-mm-ddThh:mm:ssZ”
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间,格式为“yyyy-mm-ddThh:mm:ssZ”
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowVaultSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultSummaryRequest", string(data)}, " ")
}
