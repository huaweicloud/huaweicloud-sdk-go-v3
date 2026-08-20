package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdThirdPartyAssociatedResponse Response Object
type ShowIpdThirdPartyAssociatedResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 查询失败的原因。
	Message *string `json:"message,omitempty"`

	Result         *ThirdPartyAssociatedResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowIpdThirdPartyAssociatedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdThirdPartyAssociatedResponse struct{}"
	}

	return strings.Join([]string{"ShowIpdThirdPartyAssociatedResponse", string(data)}, " ")
}
