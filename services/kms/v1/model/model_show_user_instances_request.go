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
type ShowUserInstancesRequest struct {
	VersionId string `json:"version_id"`
}

func (o ShowUserInstancesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowUserInstancesRequest", string(data)}, " ")
}
