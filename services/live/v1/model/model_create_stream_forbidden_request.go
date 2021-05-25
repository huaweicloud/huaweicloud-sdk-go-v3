package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateStreamForbiddenRequest struct {
	// op账号需要携带的特定project_id，当使用op账号时该值为所操作租户的project_id

	SpecifyProject *string `json:"specify_project,omitempty"`

	Body *StreamForbiddenSetting `json:"body,omitempty"`
}

func (o CreateStreamForbiddenRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateStreamForbiddenRequest struct{}"
	}

	return strings.Join([]string{"CreateStreamForbiddenRequest", string(data)}, " ")
}
