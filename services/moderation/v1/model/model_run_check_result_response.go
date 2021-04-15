package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunCheckResultResponse struct {
	Result         *CheckResultResultBody `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RunCheckResultResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunCheckResultResponse struct{}"
	}

	return strings.Join([]string{"RunCheckResultResponse", string(data)}, " ")
}
