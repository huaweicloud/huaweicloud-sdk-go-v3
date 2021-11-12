package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUrlAuthResponse struct{}"
	}

	return strings.Join([]string{"UpdateUrlAuthResponse", string(data)}, " ")
}
