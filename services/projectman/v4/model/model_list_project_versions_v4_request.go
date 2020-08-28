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
type ListProjectVersionsV4Request struct {
	ProjectId string `json:"project_id"`
}

func (o ListProjectVersionsV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectVersionsV4Request", string(data)}, " ")
}
