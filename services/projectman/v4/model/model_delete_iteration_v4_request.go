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
type DeleteIterationV4Request struct {
	ProjectId   string `json:"project_id"`
	IterationId int32  `json:"iteration_id"`
}

func (o DeleteIterationV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteIterationV4Request", string(data)}, " ")
}
