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
type ListResourceUsagesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListResourceUsagesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListResourceUsagesRequest", string(data)}, " ")
}
