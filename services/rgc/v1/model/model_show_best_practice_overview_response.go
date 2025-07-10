package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeOverviewResponse Response Object
type ShowBestPracticeOverviewResponse struct {

	// 总分数
	TotalScore float32 `json:"total_score,omitempty"`

	// 检测完成时间
	DetectTime *string `json:"detect_time,omitempty"`

	OrganizationAccount *BestPracticeOverviewItem `json:"organization_account,omitempty"`

	IdentityPermission *BestPracticeOverviewItem `json:"identity_permission,omitempty"`

	NetworkPlanning *BestPracticeOverviewItem `json:"network_planning,omitempty"`

	ComplianceAudit *BestPracticeOverviewItem `json:"compliance_audit,omitempty"`

	FinancialGovernance *BestPracticeOverviewItem `json:"financial_governance,omitempty"`

	DataBoundary *BestPracticeOverviewItem `json:"data_boundary,omitempty"`

	OmMonitor *BestPracticeOverviewItem `json:"om_monitor,omitempty"`

	SecurityManagement *BestPracticeOverviewItem `json:"security_management,omitempty"`
	HttpStatusCode     int                       `json:"-"`
}

func (o ShowBestPracticeOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeOverviewResponse", string(data)}, " ")
}
