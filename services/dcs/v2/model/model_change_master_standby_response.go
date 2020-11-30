/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ChangeMasterStandbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeMasterStandbyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeMasterStandbyResponse", string(data)}, " ")
}
