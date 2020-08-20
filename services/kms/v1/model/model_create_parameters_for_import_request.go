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
type CreateParametersForImportRequest struct {
	VersionId string                             `json:"version_id"`
	Body      *GetParametersForImportRequestBody `json:"body,omitempty"`
}

func (o CreateParametersForImportRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateParametersForImportRequest", string(data)}, " ")
}
