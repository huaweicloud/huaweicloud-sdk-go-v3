package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckProductHealthyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ProductInfo `json:"body,omitempty"`
}

func (o CheckProductHealthyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckProductHealthyRequest struct{}"
	}

	return strings.Join([]string{"CheckProductHealthyRequest", string(data)}, " ")
}
