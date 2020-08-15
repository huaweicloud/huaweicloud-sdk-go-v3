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
type ShowKmsTagsRequest struct {
	VersionId string `json:"version_id"`
	KeyId     string `json:"key_id"`
}

func (o ShowKmsTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowKmsTagsRequest", string(data)}, " ")
}
