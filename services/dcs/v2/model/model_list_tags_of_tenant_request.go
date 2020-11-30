/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTagsOfTenantRequest struct {
}

func (o ListTagsOfTenantRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTagsOfTenantRequest", string(data)}, " ")
}
