package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOriginalPolicyInfoResponse Response Object
type ShowOriginalPolicyInfoResponse struct {
	Policies       *Policies `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowOriginalPolicyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOriginalPolicyInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowOriginalPolicyInfoResponse", string(data)}, " ")
}
