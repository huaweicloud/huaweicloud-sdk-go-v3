/*
 * Live
 *
 * 直播服务源站所有接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateStreamForbiddenResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateStreamForbiddenResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateStreamForbiddenResponse", string(data)}, " ")
}
