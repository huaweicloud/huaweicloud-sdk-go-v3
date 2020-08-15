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
type ListGrantsRequest struct {
	VersionId string                 `json:"version_id"`
	Body      *ListGrantsRequestBody `json:"body,omitempty"`
}

func (o ListGrantsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListGrantsRequest", string(data)}, " ")
}
