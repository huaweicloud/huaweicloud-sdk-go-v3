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
type DeleteTagResponse struct {
}

func (o DeleteTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTagResponse", string(data)}, " ")
}
