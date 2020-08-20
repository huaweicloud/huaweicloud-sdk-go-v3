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

// Request Object
type CreateStreamForbiddenRequest struct {
	SpecifyProject *string                 `json:"specify_project,omitempty"`
	Body           *StreamForbiddenSetting `json:"body,omitempty"`
}

func (o CreateStreamForbiddenRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateStreamForbiddenRequest", string(data)}, " ")
}
