package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartAppResponse struct {
	// 应用id

	AppId *string `json:"app_id,omitempty"`

	State *AppState `json:"state,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAppResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartAppResponse struct{}"
	}

	return strings.Join([]string{"StartAppResponse", string(data)}, " ")
}
