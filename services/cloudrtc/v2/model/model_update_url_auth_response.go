package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateUrlAuthResponse struct {
	// 应用id

	AppId *string `json:"app_id,omitempty"`

	Authentication *AppAuth `json:"authentication,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUrlAuthResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUrlAuthResponse struct{}"
	}

	return strings.Join([]string{"UpdateUrlAuthResponse", string(data)}, " ")
}
