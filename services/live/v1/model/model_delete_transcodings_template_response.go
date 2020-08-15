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
type DeleteTranscodingsTemplateResponse struct {
}

func (o DeleteTranscodingsTemplateResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTranscodingsTemplateResponse", string(data)}, " ")
}
