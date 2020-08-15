/*
 * LiveAPI
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
type DeleteStreamForbiddenResponse struct {
}

func (o DeleteStreamForbiddenResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteStreamForbiddenResponse", string(data)}, " ")
}
