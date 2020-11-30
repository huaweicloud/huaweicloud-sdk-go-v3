/*
 * ecs
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NovaShowKeypairRequest struct {
	KeypairName         string  `json:"keypair_name"`
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`
}

func (o NovaShowKeypairRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NovaShowKeypairRequest", string(data)}, " ")
}
