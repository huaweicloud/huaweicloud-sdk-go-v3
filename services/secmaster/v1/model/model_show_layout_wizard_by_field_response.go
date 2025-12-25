package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLayoutWizardByFieldResponse Response Object
type ShowLayoutWizardByFieldResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	Data *WizardDetailInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLayoutWizardByFieldResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLayoutWizardByFieldResponse struct{}"
	}

	return strings.Join([]string{"ShowLayoutWizardByFieldResponse", string(data)}, " ")
}
