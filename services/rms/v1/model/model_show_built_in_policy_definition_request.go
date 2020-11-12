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
type ShowBuiltInPolicyDefinitionRequest struct {
	PolicyDefinitionId string `json:"policy_definition_id"`
}

func (o ShowBuiltInPolicyDefinitionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBuiltInPolicyDefinitionRequest", string(data)}, " ")
}
