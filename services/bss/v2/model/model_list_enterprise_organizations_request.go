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
type ListEnterpriseOrganizationsRequest struct {
	RecursiveQuery *int32  `json:"recursive_query,omitempty"`
	ParentId       *string `json:"parent_id,omitempty"`
}

func (o ListEnterpriseOrganizationsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListEnterpriseOrganizationsRequest", string(data)}, " ")
}
