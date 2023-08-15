package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOriginalPolicyInfoRequest Request Object
type ShowOriginalPolicyInfoRequest struct {
}

func (o ShowOriginalPolicyInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOriginalPolicyInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowOriginalPolicyInfoRequest", string(data)}, " ")
}
