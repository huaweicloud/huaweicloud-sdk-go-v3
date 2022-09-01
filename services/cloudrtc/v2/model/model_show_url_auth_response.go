package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowUrlAuthResponse struct {

	// 应用id
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	Authentication *AppAuth `json:"authentication,omitempty" xml:"authentication"`

	XRequestId     *string `json:"X-request-Id,omitempty" xml:"X-request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUrlAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUrlAuthResponse struct{}"
	}

	return strings.Join([]string{"ShowUrlAuthResponse", string(data)}, " ")
}
