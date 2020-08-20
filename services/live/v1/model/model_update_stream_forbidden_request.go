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
type UpdateStreamForbiddenRequest struct {
	SpecifyProject *string                 `json:"specify_project,omitempty"`
	Body           *StreamForbiddenSetting `json:"body,omitempty"`
}

func (o UpdateStreamForbiddenRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateStreamForbiddenRequest", string(data)}, " ")
}
