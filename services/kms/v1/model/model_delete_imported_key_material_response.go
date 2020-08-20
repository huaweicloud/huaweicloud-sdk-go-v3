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

// Response Object
type DeleteImportedKeyMaterialResponse struct {
}

func (o DeleteImportedKeyMaterialResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteImportedKeyMaterialResponse", string(data)}, " ")
}
