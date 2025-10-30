package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrustServiceStatusResponse Response Object
type ListTrustServiceStatusResponse struct {

	// xxx
	TrustedServicesEnabled *bool `json:"trusted_services_enabled,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTrustServiceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustServiceStatusResponse struct{}"
	}

	return strings.Join([]string{"ListTrustServiceStatusResponse", string(data)}, " ")
}
