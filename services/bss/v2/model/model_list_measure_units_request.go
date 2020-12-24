/*
 * BSS
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
type ListMeasureUnitsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListMeasureUnitsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListMeasureUnitsRequest", string(data)}, " ")
}
