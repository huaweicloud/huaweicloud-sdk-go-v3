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
type UpdateIterationV4Request struct {
	ProjectId   string                    `json:"project_id"`
	IterationId int32                     `json:"iteration_id"`
	Body        *CreateIterationRequestV4 `json:"body,omitempty"`
}

func (o UpdateIterationV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIterationV4Request", string(data)}, " ")
}
