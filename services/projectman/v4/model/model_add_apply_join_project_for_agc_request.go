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
type AddApplyJoinProjectForAgcRequest struct {
	DomainId  string `json:"Domain-Id"`
	UserId    string `json:"User-Id"`
	ProjectId string `json:"project_id"`
}

func (o AddApplyJoinProjectForAgcRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddApplyJoinProjectForAgcRequest", string(data)}, " ")
}
