package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalPrivacyNewResponse Response Object
type ShowGlobalPrivacyNewResponse struct {

	// 是否同意
	Policy         *bool `json:"policy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowGlobalPrivacyNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalPrivacyNewResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalPrivacyNewResponse", string(data)}, " ")
}
