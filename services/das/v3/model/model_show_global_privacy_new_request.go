package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalPrivacyNewRequest Request Object
type ShowGlobalPrivacyNewRequest struct {
}

func (o ShowGlobalPrivacyNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalPrivacyNewRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalPrivacyNewRequest", string(data)}, " ")
}
