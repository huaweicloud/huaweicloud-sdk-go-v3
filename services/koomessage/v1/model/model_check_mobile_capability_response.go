package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckMobileCapabilityResponse Response Object
type CheckMobileCapabilityResponse struct {

	// 支持智能信息手机号码列表。
	SupportedMobiles *[]string `json:"supported_mobiles,omitempty"`

	// 不支持智能信息手机号码列表。
	UnsupportedMobiles *[]string `json:"unsupported_mobiles,omitempty"`

	// 智能信息模板ID，由9位数字组成。
	TplId          *string `json:"tpl_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckMobileCapabilityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMobileCapabilityResponse struct{}"
	}

	return strings.Join([]string{"CheckMobileCapabilityResponse", string(data)}, " ")
}
