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
type ListProjectIterationsV4Request struct {
	ProjectId string `json:"project_id"`
}

func (o ListProjectIterationsV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectIterationsV4Request", string(data)}, " ")
}
