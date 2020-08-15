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
type DeleteImportedKeyMaterialRequest struct {
	VersionId string                 `json:"version_id"`
	Body      *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o DeleteImportedKeyMaterialRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteImportedKeyMaterialRequest", string(data)}, " ")
}
