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
type CreateRecordConfigResponse struct {
}

func (o CreateRecordConfigResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateRecordConfigResponse", string(data)}, " ")
}
