/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateKeyRotationIntervalRequest struct {
	VersionId string                                `json:"version_id"`
	Body      *UpdateKeyRotationIntervalRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyRotationIntervalRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateKeyRotationIntervalRequest", string(data)}, " ")
}
