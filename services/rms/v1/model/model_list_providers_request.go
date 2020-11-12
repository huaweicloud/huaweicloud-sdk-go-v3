/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProvidersRequest struct {
	Offset    *int32  `json:"offset,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListProvidersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProvidersRequest", string(data)}, " ")
}
