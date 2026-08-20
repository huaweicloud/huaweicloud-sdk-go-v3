package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpdThirdPartyResponse Response Object
type DeleteIpdThirdPartyResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 删除工作项下外部链接失败的原因。
	Message *string `json:"message,omitempty"`

	Result         *DeleteThirdPartyAssociateResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o DeleteIpdThirdPartyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpdThirdPartyResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpdThirdPartyResponse", string(data)}, " ")
}
