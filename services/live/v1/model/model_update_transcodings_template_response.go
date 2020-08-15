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
type UpdateTranscodingsTemplateResponse struct {
}

func (o UpdateTranscodingsTemplateResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateTranscodingsTemplateResponse", string(data)}, " ")
}
