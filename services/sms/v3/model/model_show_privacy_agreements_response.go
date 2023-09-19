package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivacyAgreementsResponse Response Object
type ShowPrivacyAgreementsResponse struct {

	// 查询用户是否同意隐私协议
	Flag           *bool `json:"flag,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPrivacyAgreementsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivacyAgreementsResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivacyAgreementsResponse", string(data)}, " ")
}
