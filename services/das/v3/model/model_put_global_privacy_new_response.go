package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutGlobalPrivacyNewResponse Response Object
type PutGlobalPrivacyNewResponse struct {

	// 是否同意
	Policy         *bool `json:"policy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o PutGlobalPrivacyNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutGlobalPrivacyNewResponse struct{}"
	}

	return strings.Join([]string{"PutGlobalPrivacyNewResponse", string(data)}, " ")
}
