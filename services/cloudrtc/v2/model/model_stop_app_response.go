package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopAppResponse struct {
	// 应用id

	AppId *string `json:"app_id,omitempty"`

	State *AppState `json:"state,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAppResponse struct{}"
	}

	return strings.Join([]string{"StopAppResponse", string(data)}, " ")
}
