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
type BatchCreateKmsTagsRequest struct {
	KeyId     string                         `json:"key_id"`
	VersionId string                         `json:"version_id"`
	Body      *BatchCreateKmsTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateKmsTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateKmsTagsRequest", string(data)}, " ")
}
