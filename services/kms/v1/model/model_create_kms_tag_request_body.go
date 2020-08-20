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

type CreateKmsTagRequestBody struct {
	Tag *TagItem `json:"tag,omitempty"`
}

func (o CreateKmsTagRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateKmsTagRequestBody", string(data)}, " ")
}
