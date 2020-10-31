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
type CreateIterationV4Response struct {
	// 迭代id
	Id *int32 `json:"id,omitempty"`
}

func (o CreateIterationV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateIterationV4Response", string(data)}, " ")
}
