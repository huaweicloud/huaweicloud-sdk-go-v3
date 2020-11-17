/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResetPasswordResponse struct {
}

func (o ResetPasswordResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResetPasswordResponse", string(data)}, " ")
}
