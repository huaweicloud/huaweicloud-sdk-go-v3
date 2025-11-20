package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantAssginRecordResponse Response Object
type ShowTenantAssginRecordResponse struct {

	// 租户ID
	CustomerProjectId *string `json:"customer_project_id,omitempty"`

	// 被关联租户的名称。
	CustomerName *string `json:"customer_name,omitempty"`

	// 租户的可用区。
	RegionId *string `json:"region_id,omitempty"`

	// 分配表记录总数，用于分页。
	ResourceCount *int32 `json:"resource_count,omitempty"`

	// 分配给租户的资源。
	Resources *[]AllocateSpResourceInfo `json:"resources,omitempty"`

	// 分配给租户的资源总览。
	ResourcesOverview *[]AllocateSpResourceSummaryInfo `json:"resources_overview,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTenantAssginRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantAssginRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantAssginRecordResponse", string(data)}, " ")
}
