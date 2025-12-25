package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLayoutWizardsResponse Response Object
type UpdateLayoutWizardsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 数据列表
	Data *[]WizardDetailInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLayoutWizardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLayoutWizardsResponse struct{}"
	}

	return strings.Join([]string{"UpdateLayoutWizardsResponse", string(data)}, " ")
}
