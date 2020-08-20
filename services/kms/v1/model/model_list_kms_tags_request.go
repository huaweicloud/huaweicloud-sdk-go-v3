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
type ListKmsTagsRequest struct {
	VersionId string `json:"version_id"`
}

func (o ListKmsTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKmsTagsRequest", string(data)}, " ")
}
