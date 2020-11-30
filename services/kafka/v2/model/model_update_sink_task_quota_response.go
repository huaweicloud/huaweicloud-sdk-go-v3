/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateSinkTaskQuotaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSinkTaskQuotaResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSinkTaskQuotaResponse", string(data)}, " ")
}
