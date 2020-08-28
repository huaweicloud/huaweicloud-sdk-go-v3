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
type ListProjectMembersV4Request struct {
	ProjectId string `json:"project_id"`
	Offset    *int32 `json:"offset,omitempty"`
	Limit     *int32 `json:"limit,omitempty"`
}

func (o ListProjectMembersV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectMembersV4Request", string(data)}, " ")
}
