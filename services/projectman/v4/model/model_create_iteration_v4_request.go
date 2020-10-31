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
type CreateIterationV4Request struct {
	ProjectId string                    `json:"project_id"`
	Body      *CreateIterationRequestV4 `json:"body,omitempty"`
}

func (o CreateIterationV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateIterationV4Request", string(data)}, " ")
}
