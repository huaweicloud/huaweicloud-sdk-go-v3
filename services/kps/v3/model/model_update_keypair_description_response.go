/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateKeypairDescriptionResponse struct {
}

func (o UpdateKeypairDescriptionResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateKeypairDescriptionResponse", string(data)}, " ")
}
