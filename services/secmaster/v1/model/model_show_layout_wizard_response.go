package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLayoutWizardResponse Response Object
type ShowLayoutWizardResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	Data *WizardDetailInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLayoutWizardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLayoutWizardResponse struct{}"
	}

	return strings.Join([]string{"ShowLayoutWizardResponse", string(data)}, " ")
}
