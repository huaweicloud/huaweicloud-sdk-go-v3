/*
 * AS
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
type ExecuteScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteScalingPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ExecuteScalingPolicyResponse", string(data)}, " ")
}
