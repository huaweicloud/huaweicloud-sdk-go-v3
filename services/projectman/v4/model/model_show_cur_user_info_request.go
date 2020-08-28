/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCurUserInfoRequest struct {
}

func (o ShowCurUserInfoRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCurUserInfoRequest", string(data)}, " ")
}
