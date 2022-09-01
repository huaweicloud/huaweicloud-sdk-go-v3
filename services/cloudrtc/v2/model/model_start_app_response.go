package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartAppResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	State *AppState `json:"state,omitempty" xml:"state"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAppResponse struct{}"
	}

	return strings.Join([]string{"StartAppResponse", string(data)}, " ")
}
