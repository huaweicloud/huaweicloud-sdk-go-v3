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
type UpdateKeyDescriptionRequest struct {
	VersionId string                           `json:"version_id"`
	Body      *UpdateKeyDescriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyDescriptionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateKeyDescriptionRequest", string(data)}, " ")
}
