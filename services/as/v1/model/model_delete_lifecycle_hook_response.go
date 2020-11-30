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
type DeleteLifecycleHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLifecycleHookResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteLifecycleHookResponse", string(data)}, " ")
}
