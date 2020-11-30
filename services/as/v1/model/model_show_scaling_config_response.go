/*
 * As
 *
 * 弹性伸缩API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowScalingConfigResponse struct {
	ScalingConfiguration *ScalingConfiguration `json:"scaling_configuration,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o ShowScalingConfigResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowScalingConfigResponse", string(data)}, " ")
}
