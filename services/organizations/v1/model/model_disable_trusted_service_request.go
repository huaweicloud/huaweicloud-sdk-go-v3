package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableTrustedServiceRequest Request Object
type DisableTrustedServiceRequest struct {
	Body *TrustedServiceReqBody `json:"body,omitempty"`
}

func (o DisableTrustedServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableTrustedServiceRequest struct{}"
	}

	return strings.Join([]string{"DisableTrustedServiceRequest", string(data)}, " ")
}
