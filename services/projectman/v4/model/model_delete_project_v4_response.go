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
type DeleteProjectV4Response struct {
}

func (o DeleteProjectV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteProjectV4Response", string(data)}, " ")
}
