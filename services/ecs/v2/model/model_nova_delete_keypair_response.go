/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NovaDeleteKeypairResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDeleteKeypairResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaDeleteKeypairResponse", string(data)}, " ")
}
