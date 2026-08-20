package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIpdThirdPartyResponse Response Object
type AssociateIpdThirdPartyResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 关联失败的原因。
	Message *string `json:"message,omitempty"`

	Result         *IssuesAssociationRespResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o AssociateIpdThirdPartyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIpdThirdPartyResponse struct{}"
	}

	return strings.Join([]string{"AssociateIpdThirdPartyResponse", string(data)}, " ")
}
