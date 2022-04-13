package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowUrlAuthResponse struct {
	// 应用id

	AppId *string `json:"app_id,omitempty"`

	Authentication *AppAuth `json:"authentication,omitempty"`

	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUrlAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUrlAuthResponse struct{}"
	}

	return strings.Join([]string{"ShowUrlAuthResponse", string(data)}, " ")
}
