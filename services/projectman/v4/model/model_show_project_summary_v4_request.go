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
type ShowProjectSummaryV4Request struct {
	ProjectId string `json:"project_id"`
}

func (o ShowProjectSummaryV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowProjectSummaryV4Request", string(data)}, " ")
}
