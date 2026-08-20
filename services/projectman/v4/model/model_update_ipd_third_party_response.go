package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpdThirdPartyResponse Response Object
type UpdateIpdThirdPartyResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 修改工作项下外部链接失败的原因。
	Message *string `json:"message,omitempty"`

	Result         *UpdateThirdPartyAssociateResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o UpdateIpdThirdPartyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpdThirdPartyResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpdThirdPartyResponse", string(data)}, " ")
}
