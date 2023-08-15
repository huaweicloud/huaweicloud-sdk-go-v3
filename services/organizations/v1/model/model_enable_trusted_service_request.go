package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableTrustedServiceRequest Request Object
type EnableTrustedServiceRequest struct {
	Body *TrustedServiceReqBody `json:"body,omitempty"`
}

func (o EnableTrustedServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableTrustedServiceRequest struct{}"
	}

	return strings.Join([]string{"EnableTrustedServiceRequest", string(data)}, " ")
}
