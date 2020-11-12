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
type ListBuiltInPolicyDefinitionsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListBuiltInPolicyDefinitionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBuiltInPolicyDefinitionsRequest", string(data)}, " ")
}
