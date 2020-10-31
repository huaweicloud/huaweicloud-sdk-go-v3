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

// Response Object
type AddApplyJoinProjectForAgcResponse struct {
}

func (o AddApplyJoinProjectForAgcResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddApplyJoinProjectForAgcResponse", string(data)}, " ")
}
