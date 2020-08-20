/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCustomerselfResourceRecordsRequest struct {
	Cycle               string  `json:"cycle"`
	CloudServiceType    *string `json:"cloud_service_type,omitempty"`
	Region              *string `json:"region,omitempty"`
	ChargeMode          string  `json:"charge_mode"`
	BillType            *int32  `json:"bill_type,omitempty"`
	Offset              *int32  `json:"offset,omitempty"`
	Limit               *int32  `json:"limit,omitempty"`
	ResourceId          *string `json:"resource_id,omitempty"`
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	IncludeZeroRecord   *bool   `json:"include_zero_record,omitempty"`
}

func (o ListCustomerselfResourceRecordsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCustomerselfResourceRecordsRequest", string(data)}, " ")
}
